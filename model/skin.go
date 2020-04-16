package model

import "github.com/GoodHot/TinyCMS/orm"

type Skin struct {
	orm.Model
	Dir         string `gorm:"type:varchar(60)" json:"dir"` // 模板文件夹
	Name        string `gorm:"type:varchar(60)" json:"name"`
	Description string `json:"description"`
	Version     string `gorm:"type:varchar(60)" json:"version"`
	Author      string `gorm:"type:varchar(60)" json:"author"`
	Website     string `json:"website"`
	Checksum    string `json:"checksum"`
	Active      bool   `json:"active"`
}
