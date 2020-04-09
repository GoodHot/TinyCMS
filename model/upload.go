package model

import (
	"github.com/GoodHot/TinyCMS/orm"
)

type Upload struct {
	orm.Model
	Type string `json:"type"`
	Path string `json:"path"`
	URL  string `json:"url"`
	Size int    `json:"size"`
}
