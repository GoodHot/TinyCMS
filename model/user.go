package model

import (
	"github.com/GoodHot/TinyCMS/orm"
	"time"
)

type User struct {
	orm.Model
	Nickname string `gorm:"type:varchar(60)" json:"nickname"`
	Username string `gorm:"type:varchar(60);unique_index" json:"username"`
	Password string `gorm:"type:varchar(60);not null" json:"-"`
	Avatar   string `json:"avatar"` // 头像
	Visible  bool   `json:"visible"`
	Role     int    `json:"role"`
}

func (s *User) Link() string {
	return "/author/" + s.Username
}

type UserResult struct {
	*User
	Token     string     `json:"token"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Role struct {
	orm.Model
	Name string `gorm:"type:varchar(60)" json:"name"`
}
