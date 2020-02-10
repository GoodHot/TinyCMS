package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/jinzhu/gorm"
)

type ArticleService struct {
	ChannelService *ChannelService `ioc:"auto"`
	TagService     *TagService     `ioc:"auto"`
	ORM            *orm.ORM        `ioc:"auto"`
	PageSize       int             `val:"${website.page_size}"`
	TablePrefix    string          `val:"${db.table_prefix}"`
}

func (s *ArticleService) Publish(article *model.ArticlePublish) error {
	channel, err := s.ChannelService.Get(int(article.Article.ChannelID))
	if err != nil {
		return err
	}
	tableName := s.contentTableName(channel)
	err = s.createContentTable(tableName)
	if err != nil {
		return err
	}
	return s.ORM.Tx(func(db *gorm.DB) error {
		err := db.Table(tableName).Create(article.Content).Error
		if err != nil {
			return err
		}
		article.Article.ContentTable = s.contentTableName(channel)
		article.Article.Views = 0
		article.Article.ContentID = article.Content.ID
		db.Create(article.Article)
		for _, t := range article.Tags {
			tag := s.TagService.GetByName(t.Name)
			if tag.Name == "" {
				t.ArticleCount = 0
				db.Create(t)
				tag = t
			}
			err := db.Model(&tag).UpdateColumn("article_count", gorm.Expr("article_count + ?", 1)).Error
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
	})
}

func (s *ArticleService) contentTableName(channel *model.Channel) string {
	return s.TablePrefix + "article_content_" + channel.Title
}

func (s *ArticleService) createContentTable(tableName string) error {
	if s.ORM.DB.HasTable(tableName) {
		return nil
	}
	return s.ORM.DB.Table(tableName).CreateTable(&model.ArticleContent{}).Error
}
