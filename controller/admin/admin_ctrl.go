package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/service"
	"net/http"
	"time"
)

type AdminAuthCtrl struct {
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
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()
	admin.Resume = ""
	s.AdminService.GenToken(admin)
	ctx.Put("info", admin)
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) Info(ctx *ctrl.HTTPContext) error {
	token := ctx.Context.Request().Header.Get("ACCESS-TOKEN")
	admin, err := s.AdminService.CheckToken(token)
	if err != nil || admin == nil {
		return ctx.ResultErrStatus(http.StatusUnauthorized, "please login")
	}
	ctx.Put("info", admin)
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) All(ctx *ctrl.HTTPContext) error {
	ctx.Put("admins", s.AdminService.All())
	return ctx.ResultOK()
}

type AdminForm struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
	RoleID   uint    `json:"role_id"`
	Resume   string `json:"resume"`
	Avatar   string `json:"avatar"`
}

func (s *AdminAuthCtrl) Save(ctx *ctrl.HTTPContext) error {
	form := new(AdminForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.ResultErr(err.Error())
	}
	admin := &model.Admin{
		Model: orm.Model{
			ID: form.ID,
		},
		Nickname:   form.Nickname,
		Username:   form.UserName,
		Password:   form.Password,
		Resume:     form.Resume,
		Avatar:     form.Avatar,
		RoleID:     form.RoleID,
	}
	if err := s.AdminService.Save(admin); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}
