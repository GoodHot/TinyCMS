package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ArticleService struct {
	CategoryService *CategoryService `ioc:"auto"`
	TagService      *TagService      `ioc:"auto"`
	ORM             *orm.ORM         `ioc:"auto"`
	PageSize        int              `val:"${website.page_size}"`
	TableSize       int              `val:"${website.article_table_size}"`
	TablePrefix     string           `val:"${db.table_prefix}"`
}

func (s *ArticleService) GetContent(tableName string, id int) *model.ArticleContent {
	var content model.ArticleContent
	s.ORM.DB.Table(tableName).Where("id = ?", id).First(&content)
	return &content
}

func (s *ArticleService) saveArticle(db *gorm.DB, article *model.ArticlePublish) error {
	if article.Article.Cover == "" {
		// TODO get cover
	}
	if article.Article.Description == "" {
		// TODO get description
	}
	// save content
	err := db.Table(article.Article.ContentTable).Create(article.Content).Error
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
	// update article cover,description,content_table,content_id,tags
	err = db.Model(article.Article).Updates(&model.Article{
		Description:  article.Article.Description,
		Cover:        article.Article.Cover,
		ContentTable: article.Article.ContentTable,
		ContentID:    article.Content.ID,
		Tags:         article.Article.Tags,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleService) saveTags(db *gorm.DB, article *model.ArticlePublish) error {
	db.Unscoped().Where("article_id = ?", article.Article.ID).Delete(&model.RelTagArticle{})
	if len(article.Tags) == 0 {
		return nil
	}
	for _, t := range article.Tags {
		tag := s.TagService.GetByName(t.Name)
		if tag.Name == "" {
			t.ArticleCount = 0
			db.Create(t)
			tag = t
		} else {
			t.ID = tag.ID
		}
		err := db.Model(&t).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
		if err != nil {
			return err
		}
		rel := &model.RelTagArticle{}
		rel.ArticleID = article.Article.ID
		rel.TagID = t.ID
		err = db.Create(rel).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ArticleService) Publish(article *model.ArticlePublish) error {
	s.ORM.DB.Save(article.Article)
	if article.Article.ID == 0 {
		tableName := s.contentTableName(article.Article.ID)
		err := s.createContentTable(tableName)
		if err != nil {
			return err
		}
		article.Article.ContentTable = tableName
	} else {
		article.Article = s.GetById(int(article.Article.ID))
	}
	return s.ORM.Tx(func(db *gorm.DB) error {
		return s.saveArticle(db, article)
	})
	//channel, err := s.CategoryService.Get(int(article.Article.ChannelID))
	//if err != nil {
	//	return err
	//}
	//if channel == nil {
	//	return errors.New("this channel not exists")
	//}
	//tableName := s.contentTableName(channel)
	//err = s.createContentTable(tableName)
	//if err != nil {
	//	return err
	//}
	//return s.ORM.Tx(func(db *gorm.DB) error {
	//	err := db.Table(tableName).Create(article.Content).Error
	//	if err != nil {
	//		return err
	//	}
	//	article.Article.ContentTable = s.contentTableName(channel)
	//	article.Article.Views = 0
	//	article.Article.ContentID = article.Content.ID
	//	if len(article.Tags) > 0 {
	//		tagStr := bytes.NewBufferString("")
	//		for _, t := range article.Tags {
	//			tag := s.TagService.GetByName(t.Name)
	//			if tag.Name == "" {
	//				t.ArticleCount = 0
	//				db.Create(t)
	//				tag = t
	//			}
	//			tagStr.WriteString(t.Name)
	//			tagStr.WriteString(",")
	//			err := db.Model(&tag).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
	//			if err != nil {
	//				return err
	//			}
	//			rel := &model.RelTagArticle{}
	//			rel.ArticleID = article.Article.ID
	//			rel.TagID = t.ID
	//			err = db.Create(rel).Error
	//			if err != nil {
	//				return err
	//			}
	//		}
	//		article.Article.Tags = tagStr.String()
	//	}
	//	db.Create(article.Article)
	//	return nil
	//})
}

func (s *ArticleService) Edit(article *model.ArticlePublish) error {
	return nil
	//articleDB := s.GetById(int(article.Article.ID))
	//if articleDB == nil || articleDB.Title == "" {
	//	return errors.New("article not exists")
	//}
	//s.ORM.Tx(func(db *gorm.DB) error {
	//	err := db.Table(articleDB.ContentTable).Where("id = ?", articleDB.ContentID).Updates(map[string]interface{}{
	//		"markdown": article.Content.Markdown,
	//		"html":     article.Content.Html,
	//	}).Error
	//	if err != nil {
	//		return err
	//	}
	//	tmp := article.Article
	//
	//	tags := s.TagService.GetByArticleID(int(article.Article.ID))
	//	if tags != nil && len(tags) > 0 {
	//		for _, v := range tags {
	//			err := db.Model(v).UpdateColumn("article_count", gorm.Expr("article_count - ?", 1)).Error
	//			if err != nil {
	//				return err
	//			}
	//		}
	//	}
	//	db.Unscoped().Where("article_id = ?", article.Article.ID).Delete(&model.RelTagArticle{})
	//	if len(article.Tags) > 0 {
	//		tagStr := bytes.NewBufferString("")
	//		for _, t := range article.Tags {
	//			tag := s.TagService.GetByName(t.Name)
	//			if tag.Name == "" {
	//				t.ArticleCount = 0
	//				db.Create(t)
	//				tag = t
	//			}
	//			err := db.Model(&tag).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
	//			if err != nil {
	//				return err
	//			}
	//			tagStr.WriteString(t.Name)
	//			tagStr.WriteString(",")
	//			rel := &model.RelTagArticle{}
	//			rel.ArticleID = article.Article.ID
	//			rel.TagID = tag.ID
	//			err = db.Create(rel).Error
	//			if err != nil {
	//				return err
	//			}
	//		}
	//		tmp.Tags = tagStr.String()
	//	}
	//	db.Model(tmp).Updates(map[string]interface{}{
	//		"title":        tmp.Title,
	//		"channel_id":   tmp.ChannelID,
	//		"cover":        tmp.Cover,
	//		"description":  tmp.Description,
	//		"publish_time": tmp.PublishTime,
	//		"tags":         tmp.Tags,
	//	})
	//	return nil
	//})
	//return nil
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
	return s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &[]*model.Article{},
		Where:    orm.Where(where, param...),
		OrderBy:  "id desc",
	})
}

func (s *ArticleService) Delete(id int) error {
	model := &model.Article{}
	model.ID = uint(id)
	return s.ORM.DB.Delete(model).Error
}
