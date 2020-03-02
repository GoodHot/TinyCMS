package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type RoleAuthCtrl struct {
	RoleService  *service.RoleService  `ioc:"auto"`
}

func (s *RoleAuthCtrl) AllPermission(ctx *ctrl.HTTPContext) error {
	ctx.Put("permission", s.RoleService.AllPermission())
	return ctx.ResultOK()
}

func (s *RoleAuthCtrl) InitRole(ctx *ctrl.HTTPContext) error {
	s.RoleService.Init()
	return ctx.ResultOK()
}