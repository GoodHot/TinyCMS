package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminChannelCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
}

func (s *AdminChannelCtrl) Page(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	ctx.Put("page", s.ChannelService.Page(page))
	return ctx.ResultOK()
}
