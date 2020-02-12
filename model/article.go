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
	Title        string        `json:"title"`
	ChannelID    uint          `json:"channel_id"`
	Description  string        `json:"description"`   // article description
	Cover        string        `json:"cover"`         // conver image
	ContentTable string        `json:"content_table"` // content table
	ContentID    uint          `json:"content_id"`    // content id
	PublishTime  *time.Time    `json:"publish_time"`  // timing publish
	Status       ArticleStatus `json:"status"`        // status
	Views        int           `json:"views"`         // page view count
	Tags         string        `json:"tags"`          // article tags
}

type ArticleContent struct {
	orm.Model
	Markdown string `gorm:"type:text"`
	Html     string `gorm:"type:text"`
}

type Tag struct {
	orm.Model
	Name         string `gorm:"unique_index"`
	ArticleCount int
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
