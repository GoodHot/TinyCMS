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

func (s *AdminSkinCtrl) Switch(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	if err := s.SkinService.SwitchSkin(id); err != nil {
		ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}
