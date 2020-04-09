package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminSkinCtrl struct {
	SkinService *service.SkinService `ioc:"auto"`
}

func (s *AdminSkinCtrl) All(ctx *ctrl.HTTPContext) error {
	skins, err := s.SkinService.AllSkin()
	if err != nil {
		return ctx.ResultErr(err.Error())
	}
	ctx.Put("skin", skins)
	return ctx.ResultOK()
}
