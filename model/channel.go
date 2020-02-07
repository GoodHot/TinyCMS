package model

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model
	Name         string `json:"name"`
	Title        string `json:"title"`
	SEOTitle     string `json:"seo_title"`
	Desccription string `json:"desccription"`
	Keyword      string `json:"keyword"`
	Remark       string `json:"remark"`
	Icon         string `json:"icon"`
	ParentId     uint `json:"parent_id"`
}
