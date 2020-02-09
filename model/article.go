package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title        string
	ChannelID    uint
	Description  string
	Cover        string
	ContentTable string
	ContentID    uint
}