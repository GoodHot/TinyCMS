package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminChannelCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
}

func (s *AdminChannelCtrl) List(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	s.ChannelService.List(page)
	return ctx.HTML("channel")
}
