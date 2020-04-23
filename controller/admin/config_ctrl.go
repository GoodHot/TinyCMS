package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminConfigCtrl struct {
	ConfigService *service.ConfigService `ioc:"auto"`
}

func (s *AdminConfigCtrl) ListByGroup(ctx *ctrl.HTTPContext) error {
	groupName := ctx.Param("group")
	groups := s.ConfigService.GetByGroup(groupName)
	ctx.Put("config", groups)
	return ctx.ResultOK()
}

type ConfigUpdateForm struct {
	Config []*model.Config `json:"config"`
}

func (s *AdminConfigCtrl) Update(ctx *ctrl.HTTPContext) error {
	cfg := new(ConfigUpdateForm)
	if err := ctx.Bind(cfg); err != nil {
		return ctx.ResultErr(err)
	}
	if err := s.ConfigService.Update(cfg.Config); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}
