package model


import (
	"github.com/GoodHot/TinyCMS/common/strs"
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
	CategoryID     int          `json:"category_id"`
	Description    string        `json:"description"`   // article description
	Cover          string        `json:"cover"`         // conver image
	ContentTable   string        `json:"content_table"` // content table
	ContentID      int          `json:"content_id"`    // content id
	PublishTime    *time.Time    `json:"publish_time"`  // timing publish
	Status         ArticleStatus `json:"status"`        // status
	Views          int           `json:"views"`         // page view count
	Tags           string        `json:"tags"`          // article tags
	AuthorID       int          `json:"author"`        // article author
	Template       string        `json:"template"`      // 页面渲染模板
	Visibility     uint          `json:"visibility"`    // 可见性 [1. public, 2. private]
	CreatorID      int          `json:"creator_id"`    // 创建人
	EditorID       int          `json:"editor_id"`     // 修改人
	AuthorName     string        `gorm:"-" json:"author_name"`
	CategoryName   string        `gorm:"-" json:"category_name"`
}

func (s *Article) Link() string {
	return strs.Fmt("/post/%v/%s", s.ID, s.SEOTitle)
}

type ArticleContent struct {
	orm.Model
	Markdown string `gorm:"type:text" json:"markdown"`
	Html     string `gorm:"type:text" json:"html"`
}

type Tag struct {
	orm.Model
	Name            string `gorm:"unique_index" json:"name"`
	Path            string `gorm:"unique_index" json:"path"`
	Description     string `json:"description"`
	Icon            string `json:"icon"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	ArticleCount    int    `json:"article_count"`
}

type RelTagArticle struct {
	TagID     int
	ArticleID int
}

type ArticlePublish struct {
	Article        *Article
	Tags           []*Tag
	Content        *ArticleContent
	GetCover       bool `json:"get_cover"`
	GetNetImage    bool `json:"get_net_image"`
	GetDescription bool `json:"get_description"`
}
