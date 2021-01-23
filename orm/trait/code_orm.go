package trait

import "github.com/GoodHot/TinyCMS/core"

type Code struct {
	BaseORM
	Key         string `json:"key"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Code        string `json:"code"`
}

type CodeORM interface {
	GetAll() ([]Code, *core.Err)
}
