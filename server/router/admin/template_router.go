package admin

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Template struct {
	TemplateService *service.TemplateService `ioc:"auto"`
}

func (my *Template) All(ctx *http.Context) *core.Err {
	ctx.Put("templates", my.TemplateService.AllTemplate())
	return ctx.ResultOK("ok")
}
