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