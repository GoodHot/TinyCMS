package model

import (
	"github.com/GoodHot/TinyCMS/orm"
)

type Admin struct {
	orm.Model
	Nickname   string `gorm:"type:varchar(60)" json:"nickname"`
	Username   string `gorm:"type:varchar(60);unique_index" json:"username"`
	Password   string `gorm:"size:64;not null" json:"-"`
	Background string `json:"background"`
	Resume     string `gorm:"type:varchar(140)" json:"resume"`
	Avatar     string `gorm:"type:varchar(200)" json:"avatar"`
	RoleID     uint   `json:"role_id"`
	Token      string `gorm:"-" json:"token"`
}
