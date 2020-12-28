package trait

import "github.com/GoodHot/TinyCMS/core"

type Dict struct {
	BaseORM
	Name        string   `json:"name"`
	Description string   `json:"description"`
	FormType    FormType `json:"form_type"`
	TypeValue   string   `json:"type_value"`
	Key         string   `json:"key"`
	Value       string   `json:"value"`
	Visible     bool     `json:"visible"`
}

type DictORM interface {
	GetAll() (*[]Dict, *core.Err)
	GetByKey(key string) (*Dict, *core.Err)
	UpdateValue(key string, value string) *core.Err
	Save(dict *Dict) *core.Err
}
