package model

import "github.com/GoodHot/TinyCMS/orm"

type Category struct {
	orm.Model
	Name         string      `json:"name"`
	Path         string      `json:"path"`
	SEOTitle     string      `json:"seo_title"`
	Description  string      `json:"description"`
	Keyword      string      `json:"keyword"`
	Remark       string      `json:"remark"`
	Icon         string      `json:"icon"`
	ParentId     uint        `json:"parent_id"`
	Sort         int         `json:"sort"`
	ArticleCount int         `json:"article_count"`
	Children     interface{} `gorm:"-" json:"children"`
}

func (s *Category) Link() string {
	return "/category/" + s.Path
}
