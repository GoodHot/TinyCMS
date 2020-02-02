package model

import "github.com/jinzhu/gorm"

type DictType uint

const (
	DictTypeText     = 10
	DictTypeTextArea = 20
	DictTypeRadio    = 30
	DictTypeSelect   = 40
	DictTypeSwitch   = 50
)

type Dict struct {
	gorm.Model
	Name        string
	Description string
	Type        DictType
	Options     string
	Value       string
	ReadOnly    bool
}

func (Dict) TableName() string {
	return "t_config"
}
