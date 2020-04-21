package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminAuthCtrl struct {
	AdminService *service.AdminService `ioc:"auto"`
}

type userForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AdminAuthCtrl) Login(ctx *ctrl.HTTPContext) error {
	u := new(userForm)
	if err := ctx.Bind(u); err != nil {
		return ctx.ResultErr(err)
	}
	fmt.Println(s.AdminService.EncrptyPwd("admin"))
	admin := s.AdminService.GetByUsername(u.Username)
	if admin == nil {
		return ctx.ResultErrMsg("admin not exists")
	}
	if !s.AdminService.AssertPwd(u.Password, admin.Password) {
		return ctx.ResultErrMsg("wrong password")
	}
	result := s.AdminService.GenToken(admin)
	ctx.Put("info", result)
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) All(ctx *ctrl.HTTPContext) error {
	ctx.Put("admins", s.AdminService.All(ctx.Admin.ID))
	return ctx.ResultOK()
}

//func (s *AdminAuthCtrl) Info(ctx *ctrl.HTTPContext) error {
//	token := ctx.Context.Request().Header.Get("ACCESS-TOKEN")
//	admin, err := s.AdminService.CheckToken(token)
//	if err != nil || admin == nil {
//		return ctx.ResultErrStatus(http.StatusUnauthorized, "please login")
//	}
//	ctx.Put("info", admin)
//	return ctx.ResultOK()
//}
//
//func (s *AdminAuthCtrl) All(ctx *ctrl.HTTPContext) error {
//	ctx.Put("admins", s.AdminService.All(ctx.Admin.ID))
//	return ctx.ResultOK()
//}
//
//func (s *AdminAuthCtrl) Get(ctx *ctrl.HTTPContext) error {
//	id := ctx.ParamInt("id")
//	ctx.Put("admin", s.AdminService.Get(id))
//	return ctx.ResultOK()
//}
//
//func (s *AdminAuthCtrl) Delete(ctx *ctrl.HTTPContext) error {
//	id := ctx.ParamInt("id")
//	if uint(id) == ctx.Admin.ID {
//		return ctx.ResultErr("不允许删除自己")
//	}
//	if err := s.AdminService.Delete(id); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	return ctx.ResultOK()
//}
//
//type AdminForm struct {
//	ID       uint   `json:"id"`
//	Nickname string `json:"nickname"`
//	Password string `json:"password"`
//	UserName string `json:"user_name"`
//	RoleID   uint   `json:"role_id"`
//	Resume   string `json:"resume"`
//	Avatar   string `json:"avatar"`
//}
//
//func (s *AdminAuthCtrl) Save(ctx *ctrl.HTTPContext) error {
//	form := new(AdminForm)
//	if err := ctx.Bind(form); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	admin := &model.Admin{
//		Model: orm.Model{
//			ID: form.ID,
//		},
//		Nickname: form.Nickname,
//		Username: form.UserName,
//		Password: form.Password,
//		Resume:   form.Resume,
//		Avatar:   form.Avatar,
//		RoleID:   form.RoleID,
//	}
//	if err := s.AdminService.Save(admin); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	return ctx.ResultOK()
//}
