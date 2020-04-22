package service

import (
	"github.com/GoodHot/TinyCMS/common/consts"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/go-ego/gse"
	"github.com/gomarkdown/markdown"
	"github.com/jinzhu/gorm"
	"github.com/teris-io/shortid"
	"regexp"
	"strconv"
	"strings"
)

type ArticleService struct {
	CategoryService      *CategoryService `ioc:"auto"`
	AdminService         *AdminService    `ioc:"auto"`
	TagService           *TagService      `ioc:"auto"`
	ORM                  *orm.ORM         `ioc:"auto"`
	PageSize             int              `val:"${website.page_size}"`
	TableSize            int              `val:"${website.article_table_size}"`
	TablePrefix          string           `val:"${db.table_prefix}"`
	ArticleSubTableCount int              `val:"${db.article_sub_table_count}"`
	coverRegexp          *regexp.Regexp
	descriptionRegexp    *regexp.Regexp
	tagRegexp            *regexp.Regexp
	symbolRegexp         *regexp.Regexp
	cleanTitleRegexp     *regexp.Regexp
	dict                 *gse.Segmenter
	shortIdGen           *shortid.Shortid
}

func (s *ArticleService) Startup() error {
	if s.shortIdGen == nil {
		sid, err := shortid.New(1, consts.ShortIdDefaultABC(), 9527)
		if err != nil {
			return err
		}
		s.shortIdGen = sid
	}
	return nil
}

func (s *ArticleService) GetContent(tableName string, id int) *model.ArticleContent {
	var content model.ArticleContent
	s.ORM.DB.Table(tableName).Where("id = ?", id).First(&content)
	return &content
}

func (s *ArticleService) findTitle(title string) string {
	//if s.symbolRegexp == nil {
	//	s.symbolRegexp = regexp.MustCompile("\\p{P}")
	//}
	//if s.cleanTitleRegexp == nil {
	//	s.cleanTitleRegexp = regexp.MustCompile("\\-{1,}")
	//}
	//if s.dict == nil {
	//	tmp := gse.New("zh,resource/dict/dictionary.txt", "alpha")
	//	s.dict = &tmp
	//}
	//hmm := s.dict.Cut(title, true)
	//result := bytes.NewBufferString("")
	//for i, han := range hmm {
	//	trim := strings.TrimSpace(han)
	//	if trim == "" {
	//		continue
	//	}
	//	findSymbol := s.symbolRegexp.FindAllString(trim, -1)
	//	if len(findSymbol) > 0 {
	//		trim = s.symbolRegexp.ReplaceAllString(trim, "-")
	//	}
	//	py := "" // strings.ReplaceAll(pinyin.Paragraph(trim), " ", "")
	//	result.WriteString(py)
	//	if i != len(hmm)-1 {
	//		result.WriteString("-")
	//	}
	//}
	//rst := s.cleanTitleRegexp.ReplaceAllString(result.String(), "-")
	//if rst[len(rst)-1:] == "-" {
	//	rst = rst[:len(rst)-1]
	//}
	return ""
}

func (s *ArticleService) Publish(article *model.ArticlePublish, adminID int) error {
	isInsert := article.Article.ID == 0
	var oldArticle *model.Article
	if isInsert {
		// 创建短ID
		article.Article.ShortID, _ = s.shortIdGen.Generate()
		article.Article.CreatorID = adminID
	} else {
		oldArticle = s.GetById(article.Article.ID)
	}
	updateMap := make(map[string]interface{})
	// 0. 保存article
	if isInsert {
		s.ORM.DB.Save(article.Article)
	} else {
		updateMap["title"] = article.Article.Title
		updateMap["seo_title"] = article.Article.SEOTitle
		updateMap["seo_description"] = article.Article.SEODescription
		updateMap["seo_keyword"] = article.Article.SEOKeyword
		updateMap["category_id"] = article.Article.CategoryID
		updateMap["description"] = article.Article.Description
		updateMap["cover"] = article.Article.Cover
		updateMap["status"] = article.Article.Status
		updateMap["author"] = article.Article.AuthorID
		updateMap["template"] = article.Article.Template
		updateMap["visible"] = article.Article.Visible
		updateMap["editor_id"] = adminID
	}
	// 1. 保存content
	contentTable := s.ORM.ArticleContentTableName(article.Article.ID % s.ArticleSubTableCount)
	// TODO 这里默认把markdown转成html
	md := []byte(article.Content.Source)
	article.Content.Html = string(markdown.ToHTML(md, nil, nil))
	// 暂时这样处理 👆
	if isInsert {
		s.ORM.DB.Table(contentTable).Save(article.Content)
		article.Article.ContentTable = contentTable
		article.Article.ContentID = article.Content.ID
		updateMap["content_id"] = article.Content.ID
		updateMap["content_table"] = article.Article.ContentTable
	} else {
		contentUpdate := make(map[string]interface{})
		contentUpdate["html"] = article.Content.Html
		contentUpdate["source"] = article.Content.Source
		s.ORM.DB.Table(oldArticle.ContentTable).Where("id = ?", oldArticle.ContentID).Updates(contentUpdate)
	}
	// 2. 保存tag和tag rel
	// 如果不是新插入，则先删除原来的tag关联
	// TODO 这里有性能问题，应该先对比新老标签，取出交集然后再更新和删除
	if !isInsert {
		tags := s.TagService.GetByArticleID(oldArticle.ID)
		if len(tags) > 0 {
			s.ORM.DB.Unscoped().Where("article_id = ?", oldArticle.ID).Delete(&model.RelTagArticle{})
			var ids []int
			for _, tag := range tags {
				ids = append(ids, tag.ID)
			}
			s.ORM.DB.Model(&model.Tag{}).Where("id in (?)", ids).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1))
		}
	}
	if len(article.Tags) > 0 {
		var tagNames []string
		for _, tag := range article.Tags {
			s.saveTag(tag, 0)
			s.ORM.DB.Model(&model.Tag{}).Where("id = ?", tag.ID).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1))
			s.ORM.DB.Save(&model.RelTagArticle{
				TagID:     tag.ID,
				ArticleID: article.Article.ID,
			})
			tagNames = append(tagNames, tag.Name)
		}
		updateMap["tags"] = strings.Join(tagNames, ",")
	}
	// 3. 保存category
	if article.Article.CategoryID != 0 {
		if oldArticle != nil && oldArticle.CategoryID != article.Article.CategoryID {
			s.ORM.DB.Model(&model.Category{}).Where("id = ?", oldArticle.CategoryID).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1))
		}
		s.ORM.DB.Model(&model.Category{}).Where("id = ?", article.Article.CategoryID).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1))
	}
	// 4. 更新article
	return s.ORM.DB.Model(&model.Article{}).Where("id = ?", article.Article.ID).Updates(updateMap).Error
}

func (s *ArticleService) saveTag(tag *model.Tag, index int) {
	if tag.Path == "" {
		pathAppend := ""
		if index != 0 {
			pathAppend = strconv.Itoa(index)
		}
		tag.Path = tag.Name + pathAppend
	}
	var db model.Tag
	s.ORM.DB.Model(&model.Tag{}).Where("name = ?", tag.Name).First(&db)
	if db.Name != "" {
		tag.ID = db.ID
		return
	}
	s.ORM.DB.Model(&model.Tag{}).Where("path = ?", tag.Path).First(&db)
	if db.Name != "" {
		s.saveTag(tag, index+1)
		return
	}
	s.ORM.DB.Save(&tag)
}

func (s *ArticleService) GetById(id int) *model.Article {
	var article model.Article
	s.ORM.DB.Where("id = ?", id).Find(&article)
	return &article
}

func (s *ArticleService) GetByShortId(shortiD string) *model.Article {
	var article model.Article
	s.ORM.DB.Where("short_id = ?", shortiD).Find(&article)
	return &article
}

type ArticlePageQuery struct {
	Type       int
	CategoryID int
	Keyword    string
	TagID      int
	UserID     int
	Status     int
	Visible    int
}

func (s *ArticleService) Page(page int, query *ArticlePageQuery) *orm.Page {
	where := "1 = 1"
	var param []interface{}
	if query.CategoryID != 0 {
		if query.CategoryID == -1 {
			query.CategoryID = 0
		}
		where += " and category_id = ?"
		param = append(param, query.CategoryID)
	}
	if query.Type != 0 {
		where += " and status = ?"
		param = append(param, query.Type)
	}
	if query.UserID != 0 {
		where += " and author_id = ?"
		param = append(param, query.UserID)
	}
	query.Keyword = strings.TrimSpace(query.Keyword)
	if query.Keyword != "" {
		where += " and (title like ? or description like ?)"
		param = append(param, "%"+query.Keyword+"%")
		param = append(param, "%"+query.Keyword+"%")
	}
	if query.TagID != 0 {
		articleIDs := s.TagService.getArticleIDs(query.TagID)
		where += " and id in (?)"
		param = append(param, articleIDs)
	}
	if query.Status != 0 {
		where += " and status = ?"
		param = append(param, query.Status)
	}
	if query.Visible != 0 {
		where += " and visible = ?"
		param = append(param, query.Visible)
	}
	var article []*model.Article
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &article,
		Where:    orm.Where(where, param...),
		OrderBy:  "status desc, updated_at desc",
	})
	return result
}

func (s *ArticleService) Delete(ids []int) error {
	return s.ORM.Tx(func(db *gorm.DB) error {
		for _, id := range ids {
			article := s.GetById(id)
			if article == nil || article.ID == 0 {
				continue
			}
			err := db.Where("id = ?", id).Delete(&model.Article{}).Error
			if err != nil {
				return err
			}
			tags := s.TagService.GetByArticleID(id)
			if tags != nil && len(tags) > 0 {
				for _, tag := range tags {
					err := db.Model(&model.Tag{}).Where("id in (?)", tag.ID).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1)).Error
					if err != nil {
						return err
					}
				}
			}
			err = db.Table(article.ContentTable).Where("id = ?", article.ContentID).Delete(&model.ArticleContent{}).Error
			if err != nil {
				return err
			}
			err = db.Unscoped().Where("article_id = ?", id).Delete(&model.RelTagArticle{}).Error
			if err != nil {
				return err
			}
			err = db.Model(&model.Category{}).Where("id = ?", article.CategoryID).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1)).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
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

func (s *ArticleService) TotalCount() int {
	count := 0
	s.ORM.DB.Model(&model.Article{}).Count(&count)
	return count
}
