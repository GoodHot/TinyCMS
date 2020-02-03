package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminChannelCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
}

func (*AdminChannelCtrl) List(ctx *ctrl.HTTPContext) error {
	return ctx.ResultOK()
}
