package admin

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type CodeInject struct {
	CodeService *service.CodeService `ioc:"auto"`
}

func (my *CodeInject) All(ctx *http.Context) *core.Err {
	codes, err := my.CodeService.GetAll()
	if err != nil {
		return err
	}
	ctx.Put("codes", codes)
	return ctx.ResultOK("ok")
}

func (my *CodeInject) Save(ctx *http.Context) *core.Err {
	var code trait.Code
	if err := ctx.Bind(&code); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	if err := my.CodeService.Save(&code); err != nil {
		return err
	}
	ctx.Put("id", code.ID)
	return ctx.ResultOK("ok")
}

func (my *CodeInject) Delete(ctx *http.Context) *core.Err {
	id := ctx.ParamInt("id")
	if id == 0 {
		return core.NewErr(core.Err_Sys_Server)
	}
	if err := my.CodeService.DeleteByID(id); err != nil {
		return err
	}
	return ctx.ResultOK("ok")
}
