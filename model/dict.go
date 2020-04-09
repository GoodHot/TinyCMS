package model

import (
	"github.com/GoodHot/TinyCMS/orm"
)

type DictType string

const (
	DictTypeText        = "text"
	DictTypeTextArea    = "textarea"
	DictTypeSelect      = "select" // 格式为 key:value,key:value
	DictTypeUploadImage = "uploadImage"
)

type Dict struct {
	orm.Model
	Key         string   `json:"key"`         // 唯一Key
	Name        string   `json:"name"`        // 名称
	Description string   `json:"description"` // 描述
	Type        DictType `json:"type"`        // 数据类型
	Options     string   `json:"options"`     // 选项
	Value       string   `json:"value"`       // 值
	ReadOnly    bool     `json:"read_only"`   // 是否只读
	Visible     bool     `json:"visible"`     // 是否可见
}
