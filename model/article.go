package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title     string
	ChannelID uint
}

func (Article) TableName() string {
	return "t_article"
}

//type ArticleData struct {
//	gorm.Model
//	ArticleID uint
//}
//
//func (ad ArticleData) TableName() string {
//	return "t_article_" + strconv.Itoa(int(ad.ChannelID))
//}
