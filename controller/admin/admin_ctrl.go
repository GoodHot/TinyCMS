package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminAuthCtrl struct {
	RoleService  *service.RoleService  `ioc:"auto"`
	AdminService *service.AdminService `ioc:"auto"`
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AdminAuthCtrl) Login(ctx *ctrl.HTTPContext) error {
	u := new(user)
	if err := ctx.Bind(u); err != nil {
		return ctx.ResultErr(err.Error())
	}
	admin := s.AdminService.GetByUsername(u.Username)
	if admin == nil {
		return ctx.ResultErr("admin not exists")
	}
	if !s.AdminService.AssertPwd(u.Password, admin.Password) {
		return ctx.ResultErr("wrong password")
	}
	s.AdminService.GenToken(admin)
	ctx.Put("info", admin)
	return ctx.ResultOK()
}

func (*AdminAuthCtrl) Info(ctx *ctrl.HTTPContext) error {
	admin := model.Admin{}
	admin.Avatar = "http://wx2.sinaimg.cn/mw600/6b465dably1g9jos0m98ej20au0nsgmp.jpg"
	admin.Token = "test"
	ctx.Put("info", admin)
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) InitRole(ctx *ctrl.HTTPContext) error {
	s.RoleService.Init()
	return ctx.ResultOK()
}
