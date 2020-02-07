package model

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"type:varchar(60);unique_index"`
	Password string `gorm:"size:64;not null"`
	Avatar   string `json:"avatar"`
	Token    string `gorm:"-" json:"token"`
}
