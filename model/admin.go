package model

import (
	"github.com/GoodHot/TinyCMS/orm"
	"time"
)

type Admin struct {
	orm.Model
	Nickname string `gorm:"type:varchar(60)" json:"nickname"`
	Username string `gorm:"type:varchar(60);unique_index" json:"username"`
	Password string `gorm:"type:varchar(60);not null" json:"-"`
}

type AdminResult struct {
	*Admin
	Token     string     `json:"token"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}


type Admin333 struct {
	orm.Model
	Nickname string `gorm:"type:varchar(60)" json:"nickname"`
	Username string `gorm:"type:varchar(60);unique_index" json:"username"`
	Password string `gorm:"type:varchar(60);not null" json:"-"`
}