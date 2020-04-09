package model

import (
	"github.com/GoodHot/TinyCMS/orm"
)

type DictType string

const (
	DictTypeText     = "text"
	DictTypeTextArea = "textarea"
	DictTypeSelect   = "select"
)

type Dict struct {
	orm.Model
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        DictType `json:"type"`
	Options     string   `json:"options"`
	Value       string   `json:"value"`
	ReadOnly    bool     `json:"read_only"`
	Visible     bool     `json:"visible"`
}
