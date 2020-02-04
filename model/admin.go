package model

import "github.com/jinzhu/gorm"

type AdminModel struct {
	gorm.Model
	Username string `gorm:"type:varchar(60);unique_index"`
	Password string `gorm:"size:64"`
}
