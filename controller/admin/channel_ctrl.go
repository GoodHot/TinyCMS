package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminChannelCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
}

func (s *AdminChannelCtrl) Page(ctx *ctrl.HTTPContext) error {
	parentId := ctx.ParamInt("parentId")
	page := ctx.QueryParamInt("page")
	ctx.Put("page", s.ChannelService.Page(page, parentId))
	return ctx.ResultOK()
}

func (s *AdminChannelCtrl) Save(ctx *ctrl.HTTPContext) error {
	channel := new(model.Channel)
	if err := ctx.Bind(channel); err != nil {
		return ctx.ResultErr(err.Error())
	}
	if err := s.ChannelService.Save(channel); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}

func (s *AdminChannelCtrl) Delete(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	if err := s.ChannelService.Delete(id); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}

func (s *AdminChannelCtrl) Get(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	channel, err := s.ChannelService.Get(id)
	if err != nil {
		return ctx.ResultErr(err.Error())
	}
	ctx.Put("channel", channel)
	return ctx.ResultOK()
}
