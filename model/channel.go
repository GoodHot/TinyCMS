package model

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model
	Name         string
	Title        string
	Desccription string
	ParentId     uint
}

func (Channel) TableName() string {
	return "t_channel"
}
