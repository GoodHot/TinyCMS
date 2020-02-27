package service

import (
	"bytes"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/go-ego/gse"
	"github.com/jinzhu/gorm"
	"github.com/mozillazg/go-pinyin"
	"regexp"
	"strconv"
	"strings"
)

type ArticleService struct {
	CategoryService   *CategoryService `ioc:"auto"`
	AdminService      *AdminService    `ioc:"auto"`
	TagService        *TagService      `ioc:"auto"`
	ORM               *orm.ORM         `ioc:"auto"`
	PageSize          int              `val:"${website.page_size}"`
	TableSize         int              `val:"${website.article_table_size}"`
	TablePrefix       string           `val:"${db.table_prefix}"`
	coverRegexp       *regexp.Regexp
	descriptionRegexp *regexp.Regexp
	tagRegexp         *regexp.Regexp
	symbolRegexp      *regexp.Regexp
	cleanTitleRegexp  *regexp.Regexp
	dict              *gse.Segmenter
}

func (s *ArticleService) GetContent(tableName string, id int) *model.ArticleContent {
	var content model.ArticleContent
	s.ORM.DB.Table(tableName).Where("id = ?", id).First(&content)
	return &content
}

func (s *ArticleService) findCover(html string) string {
	if s.coverRegexp == nil {
		s.coverRegexp = regexp.MustCompile("<img\\s+src=\"(.+?)\"\\s+.+?/>")
	}
	match := s.coverRegexp.FindAllStringSubmatch(html, -1)
	if match == nil || len(match) == 0 {
		return ""
	}
	return match[0][1]
}

func (s *ArticleService) findDescription(html string) string {
	if s.descriptionRegexp == nil {
		s.descriptionRegexp = regexp.MustCompile("<p>(.+?)</p>")
	}
	if s.tagRegexp == nil {
		s.tagRegexp = regexp.MustCompile("(<.+?>.+?</.+?>|<.+?/>)")
	}
	submatch := s.descriptionRegexp.FindAllStringSubmatch(html, -1)
	for _, m := range submatch {
		s := s.tagRegexp.ReplaceAllString(m[1], "")
		s = strings.TrimSpace(s)
		if s != "" {
			temp := []rune(s)
			if len(temp) > 150 {
				return string(temp[:150])
			}
			return string(temp)
		}
	}
	return ""
}

func (s *ArticleService) findTitle(title string) string {
	if s.symbolRegexp == nil {
		s.symbolRegexp = regexp.MustCompile("\\p{P}")
	}
	if s.cleanTitleRegexp == nil {
		s.cleanTitleRegexp = regexp.MustCompile("\\-{1,}")
	}
	if s.dict == nil {
		tmp := gse.New("zh,resource/dict/dictionary.txt", "alpha")
		s.dict = &tmp
	}
	hmm := s.dict.Cut(title, true)
	result := bytes.NewBufferString("")
	for i, han := range hmm {
		trim := strings.TrimSpace(han)
		if trim == "" {
			continue
		}
		findSymbol := s.symbolRegexp.FindAllString(trim, -1)
		if len(findSymbol) > 0 {
			trim = s.symbolRegexp.ReplaceAllString(trim, "-")
		}
		py := strings.ReplaceAll(pinyin.Paragraph(trim), " ", "")
		result.WriteString(py)
		if i != len(hmm)-1 {
			result.WriteString("-")
		}
	}
	rst := s.cleanTitleRegexp.ReplaceAllString(result.String(), "-")
	if rst[len(rst)-1:] == "-" {
		rst = rst[:len(rst)-1]
	}
	return rst
}

func (s *ArticleService) saveArticle(db *gorm.DB, article *model.ArticlePublish) error {
	dbArticle := &model.Article{}
	db.Where("id = ?", article.Article.ID).Find(dbArticle)
	if dbArticle.CategoryID != 0 && dbArticle.CategoryID != article.Article.CategoryID {
		err := db.Model(&model.Category{}).Where("id = ?", dbArticle.CategoryID).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1)).Error
		if err != nil {
			return err
		}
	}
	err := db.Model(&model.Category{}).Where("id = ?", article.Article.CategoryID).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
	if err != nil {
		return err
	}
	article.Article.Cover = strings.TrimSpace(article.Article.Cover)
	if article.Article.Cover == "" {
		article.Article.Cover = s.findCover(article.Content.Html)
	}
	article.Article.Description = strings.TrimSpace(article.Article.Description)
	if article.Article.Description == "" {
		article.Article.Description = s.findDescription(article.Content.Html)
	}
	article.Article.SEOTitle = strings.TrimSpace(article.Article.SEOTitle)
	if article.Article.SEOTitle == "" {
		article.Article.SEOTitle = s.findTitle(article.Article.Title)
	}
	if dbArticle.ContentID != 0 {
		dbContent := &model.ArticleContent{}
		db.Table(article.Article.ContentTable).Where("id = ?", dbArticle.ContentID).Find(dbContent)
		article.Content.ID = dbContent.ID
	}
	// save content
	err = db.Table(article.Article.ContentTable).Save(article.Content).Error
	if err != nil {
		return err
	}
	// save tags
	err = s.saveTags(db, article)
	if err != nil {
		return err
	}
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			article.Article.Tags += tag.Name + ","
		}
	}
	// update article cover,description,content_table,content_id,tags, seo_title
	err = db.Model(article.Article).Updates(&model.Article{
		Description:  article.Article.Description,
		Cover:        article.Article.Cover,
		ContentTable: article.Article.ContentTable,
		ContentID:    article.Content.ID,
		Tags:         article.Article.Tags,
		SEOTitle:     article.Article.SEOTitle,
		CategoryID:   article.Article.CategoryID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleService) saveTags(db *gorm.DB, article *model.ArticlePublish) error {
	// 1. 减去t_tag表中article_count的数量 (article_count - 1)
	var rels []*model.RelTagArticle
	db.Where("article_id = ?", article.Article.ID).Find(&rels)
	if len(rels) > 0 {
		var ids []uint
		for _, t := range rels {
			ids = append(ids, t.TagID)
		}
		err := db.Model(&model.Tag{}).Where("id in (?)", ids).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1)).Error
		if err != nil {
			return err
		}
	}
	// 2. 删除t_rel_tag_article 关联关系
	db.Unscoped().Where("article_id = ?", article.Article.ID).Delete(&model.RelTagArticle{})
	if len(article.Tags) == 0 {
		return nil
	}
	// 3. 重新增加新的关联关系
	var ids []uint
	for _, t := range article.Tags {
		tag := s.TagService.GetByName(t.Name)
		if tag.Name == "" {
			t.ArticleCount = 0
			db.Create(t)
			tag = t
			s.TagService.PutPrefix(tag)
		} else {
			t.ID = tag.ID
		}
		ids = append(ids, t.ID)
		rel := &model.RelTagArticle{}
		rel.ArticleID = article.Article.ID
		rel.TagID = t.ID
		err := db.Create(rel).Error
		if err != nil {
			return err
		}
	}
	// 4. t_tag表中article_count + 1
	err := db.Model(&model.Tag{}).Where("id in (?)", ids).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleService) Publish(article *model.ArticlePublish) error {
	if article.Article.ID == 0 {
		s.ORM.DB.Save(article.Article)
	} else {
		s.ORM.DB.Model(article.Article).Updates(&model.Article{
			Title:  article.Article.Title,
			Status: article.Article.Status,
		})
	}
	tableName := s.contentTableName(article.Article.ID)
	err := s.createContentTable(tableName)
	if err != nil {
		return err
	}
	article.Article.ContentTable = tableName
	return s.ORM.Tx(func(db *gorm.DB) error {
		return s.saveArticle(db, article)
	})
}

func (s *ArticleService) Edit(article *model.ArticlePublish) error {
	return nil
}

func (s *ArticleService) GetById(id int) *model.Article {
	var article model.Article
	s.ORM.DB.Where("id = ?", id).Find(&article)
	return &article
}

func (s *ArticleService) contentTableName(articleID uint) string {
	mod := int(articleID) % s.TableSize
	return s.TablePrefix + "article_content_" + strconv.Itoa(mod)
}

func (s *ArticleService) createContentTable(tableName string) error {
	if s.ORM.DB.HasTable(tableName) {
		return nil
	}
	return s.ORM.DB.Table(tableName).CreateTable(&model.ArticleContent{}).Error
}

type ArticlePageQuery struct {
	ChannelID int
}

func (s *ArticleService) Page(page int, query ArticlePageQuery) *orm.Page {
	where := "1 = 1"
	var param []interface{}
	if query.ChannelID != 0 {
		where += " and channel_id = ?"
		param = append(param, query.ChannelID)
	}
	var article []*model.Article
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &article,
		Where:    orm.Where(where, param...),
		OrderBy:  "status desc, id desc",
	})
	if len(article) > 0 {
		for _, v := range article {
			admin := s.AdminService.TrieGet(v.AuthorID)
			if admin != nil {
				v.AuthorName = admin.Nickname
			}
			category := s.CategoryService.TrieGet(v.CategoryID)
			if category != nil {
				v.CategoryName = category.Name
			} else {
				v.CategoryName = "未归类"
			}
		}
	}
	return result
}

func (s *ArticleService) Delete(id int) error {
	model := &model.Article{}
	model.ID = uint(id)
	return s.ORM.DB.Delete(model).Error
}

func (s *ArticleService) NoCategoryCount() int {
	count := 0
	s.ORM.DB.Model(&model.Article{}).Where("category_id = 0").Count(&count)
	return count
}

func (s *ArticleService) DraftCount() int {
	count := 0
	s.ORM.DB.Model(&model.Article{}).Where("status = ?", model.ArticleStatusDraft).Count(&count)
	return count
}
