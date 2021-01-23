package admin

import (
	"github.com/GoodHot/TinyCMS/core"
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
