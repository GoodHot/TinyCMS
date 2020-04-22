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
	ShortID        string        `gorm:"UNIQUE_INDEX" json:"short_id"`
	Title          string        `json:"title"`
	SEOTitle       string        `json:"seo_title"`
	SEODescription string        `json:"seo_description"`
	SEOKeyword     string        `json:"seo_keyword"`
	CategoryID     int           `json:"category_id"`
	Description    string        `json:"description"`   // article description
	Cover          string        `json:"cover"`         // conver image
	ContentTable   string        `json:"content_table"` // content table
	ContentID      int           `json:"content_id"`    // content id
	PublishTime    *time.Time    `json:"publish_time"`  // timing publish
	Status         ArticleStatus `json:"status"`        // status
	Views          int           `json:"views"`         // page view count
	Tags           string        `json:"tags"`          // article tags
	AuthorID       int           `json:"author"`        // article author
	Template       string        `json:"template"`      // 页面渲染模板
	Visible        uint          `json:"visible"`       // 可见性 [1. public, 2. private]
	CreatorID      int           `json:"creator_id"`    // 创建人
	EditorID       int           `json:"editor_id"`     // 修改人
}

func (s *Article) Link() string {
	if s.SEOTitle == "" {
		return strs.Fmt("/post/%v", s.ID)
	}
	return strs.Fmt("/post/%v/%s", s.ID, s.SEOTitle)
}

type ArticleContent struct {
	orm.Model
	Source string `gorm:"type:text" json:"source"`
	Html   string `gorm:"type:text" json:"html"`
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
	Article        *Article        `json:"article"`
	Tags           []*Tag          `json:"tags"`
	Content        *ArticleContent `json:"content"`
	GetCover       bool            `json:"get_cover"`
	GetNetImage    bool            `json:"get_net_image"`
	GetDescription bool            `json:"get_description"`
}
