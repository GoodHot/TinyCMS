package model

import "github.com/GoodHot/TinyCMS/orm"

type Channel struct {
	orm.Model
	Name         string      `json:"name"`
	Title        string      `json:"title"`
	SEOTitle     string      `json:"seo_title"`
	Desccription string      `json:"desccription"`
	Keyword      string      `json:"keyword"`
	Remark       string      `json:"remark"`
	Icon         string      `json:"icon"`
	ParentId     uint        `json:"parent_id"`
	HasChild     bool        `json:"has_child"`
	Sort         int         `json:"sort"`
	Children     interface{} `gorm:"-" json:"children"`
}
