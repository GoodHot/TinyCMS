package model

import (
	"github.com/GoodHot/TinyCMS/orm"
	"time"
)

type ArticleStatus int

const (
	ArticleStatusPublish ArticleStatus = 1
	ArticleStatusDraft   ArticleStatus = 2
)

type Article struct {
	orm.Model
	Title          string        `json:"title"`
	SEOTitle       string        `json:"seo_title"`
	SEODescription string        `json:"seo_description"`
	SEOKeyword     string        `json:"seo_keyword"`
	CategoryID     uint          `json:"category_id"`
	Description    string        `json:"description"`   // article description
	Cover          string        `json:"cover"`         // conver image
	ContentTable   string        `json:"content_table"` // content table
	ContentID      uint          `json:"content_id"`    // content id
	PublishTime    *time.Time    `json:"publish_time"`  // timing publish
	Status         ArticleStatus `json:"status"`        // status
	Views          int           `json:"views"`         // page view count
	Tags           string        `json:"tags"`          // article tags
	AuthorID       uint          `json:"author"`        // article author
	Template       uint          `json:"template"`      // 页面渲染模板
	Visibility     uint          `json:"visibility"`    // 可见性 [1. public, 2. private]
	CreatorId      uint          `json:"creator_id"`    // 创建人
	Editor         uint          `json:"editor"`        // 修改人
	AuthorName     string        `gorm:"-" json:"author_name"`
	CategoryName   string        `gorm:"-" json:"category_name"`
}

type ArticleContent struct {
	orm.Model
	Markdown string `gorm:"type:text" json:"markdown"`
	Html     string `gorm:"type:text" json:"html"`
}

type Tag struct {
	orm.Model
	Name         string `gorm:"unique_index" json:"name"`
	ArticleCount int    `json:"article_count"`
}

type RelTagArticle struct {
	TagID     uint
	ArticleID uint
}

type ArticlePublish struct {
	Article *Article
	Tags    []*Tag
	Content *ArticleContent
}
