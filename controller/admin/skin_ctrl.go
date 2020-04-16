package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminSkinCtrl struct {
	SkinService *service.SkinService `ioc:"auto"`
}

func (s *AdminSkinCtrl) List(ctx *ctrl.HTTPContext) error {
	return ctx.ResultOK()
}

func (s *AdminSkinCtrl) Install(ctx *ctrl.HTTPContext) error {
	return ctx.ResultOK()
}
