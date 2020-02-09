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
	Title        string
	ChannelID    uint
	Description  string        // article description
	Cover        string        // conver image
	ContentTable string        // content table
	ContentID    uint          // content id
	PublishTime  *time.Time    // timing publish
	Status       ArticleStatus // status
}

type ArticleContent struct {
	orm.Model
	ArticleID uint
	Markdown  string `gorm:"type:text"`
	Html      string `gorm:"type:text"`
}

type Tag struct {
	orm.Model
	Name         string
	ArticleCount int
}

type RelTagArticle struct {
	TagID     uint
	ArticleID uint
}
