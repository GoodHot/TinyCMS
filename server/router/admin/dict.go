package admin

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Dict struct {
	DictService *service.DictService `ioc:"auto"`
}

type dictUpdate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (my *Dict) Update(ctx *http.Context) *core.Err {
	var dict dictUpdate
	if err := ctx.Bind(&dict); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	if err := my.DictService.Update(dict.Key, dict.Value); err != nil {
		return err
	}
	return ctx.ResultOK("ok")
}

func (my *Dict) Get(ctx *http.Context) *core.Err {
	key := ctx.Param("key")
	dict, err := my.DictService.Get(key)
	if err != nil {
		return err
	}
	ctx.Put("dict", dict)
	return ctx.ResultOK("ok")
}
