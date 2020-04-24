package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminAuthCtrl struct {
	UserService *service.UserService `ioc:"auto"`
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
	admin := s.UserService.GetByUsername(u.Username)
	if admin == nil {
		return ctx.ResultErrMsg("admin not exists")
	}
	if !s.UserService.AssertPwd(u.Password, admin.Password) {
		return ctx.ResultErrMsg("wrong password")
	}
	result := s.UserService.GenToken(admin)
	ctx.Put("info", result)
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) All(ctx *ctrl.HTTPContext) error {
	ctx.Put("users", s.UserService.All(ctx.User.ID))
	return ctx.ResultOK()
}

type UserCreateForm struct {
	model.User
	Password string `json:"password"`
}

func (s *AdminAuthCtrl) Save(ctx *ctrl.HTTPContext) error {
	form := new(UserCreateForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.ResultErr(err)
	}
	user := &model.User{
		Model: orm.Model{
			ID: form.ID,
		},
		Nickname: form.Nickname,
		Username: form.Username,
		Password: form.Password,
		Avatar:   form.Avatar,
		Visible:  form.Visible,
		Role:     0,
	}
	if err := s.UserService.Save(user); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}

type UpdatePwdForm struct {
	Password string
}

func (s *AdminAuthCtrl) Password(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	form := new(UpdatePwdForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.ResultErr(err)
	}
	if err := s.UserService.UpdatePwd(id, form.Password); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) Get(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	ctx.Put("user", s.UserService.GetByID(id))
	return ctx.ResultOK()
}

func (s *AdminAuthCtrl) InitAdmin(ctx *ctrl.HTTPContext) error {
	s.UserService.Save(&model.User{
		Nickname: "admin",
		Username: "admin",
		Password: "admin",
	})
	return ctx.ResultOK()
}

//func (s *AdminAuthCtrl) Info(ctx *ctrl.HTTPContext) error {
//	token := ctx.Context.Request().Header.Get("ACCESS-TOKEN")
//	admin, err := s.UserService.CheckToken(token)
//	if err != nil || admin == nil {
//		return ctx.ResultErrStatus(http.StatusUnauthorized, "please login")
//	}
//	ctx.Put("info", admin)
//	return ctx.ResultOK()
//}
//
//func (s *AdminAuthCtrl) All(ctx *ctrl.HTTPContext) error {
//	ctx.Put("admins", s.UserService.All(ctx.Admin.ID))
//	return ctx.ResultOK()
//}
//
//func (s *AdminAuthCtrl) Get(ctx *ctrl.HTTPContext) error {
//	id := ctx.ParamInt("id")
//	ctx.Put("admin", s.UserService.Get(id))
//	return ctx.ResultOK()
//}
//
//func (s *AdminAuthCtrl) Delete(ctx *ctrl.HTTPContext) error {
//	id := ctx.ParamInt("id")
//	if uint(id) == ctx.Admin.ID {
//		return ctx.ResultErr("不允许删除自己")
//	}
//	if err := s.UserService.Delete(id); err != nil {
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
//	if err := s.UserService.Save(admin); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	return ctx.ResultOK()
//}
