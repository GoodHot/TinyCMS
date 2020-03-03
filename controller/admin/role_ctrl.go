package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/service"
)

type RoleCtrl struct {
	RoleService *service.RoleService `ioc:"auto"`
}

func (s *RoleCtrl) AllPermission(ctx *ctrl.HTTPContext) error {
	ctx.Put("permission", s.RoleService.AllPermission())
	return ctx.ResultOK()
}

func (s *RoleCtrl) InitRole(ctx *ctrl.HTTPContext) error {
	s.RoleService.Init()
	return ctx.ResultOK()
}

type RoleForm struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Permission []uint `json:"permission"`
}

func (s *RoleCtrl) Save(ctx *ctrl.HTTPContext) error {
	form := new(RoleForm)
	if err := ctx.Bind(form); err != nil {
		return ctx.ResultErr(err.Error())
	}
	if form.Permission == nil || len(form.Permission) == 0 {
		return ctx.ResultErr("请选择角色拥有权限")
	}
	role := model.Role{
		Code:     form.Code,
		RoleName: form.Name,
		IsSuper:  false,
	}
	var permission []*model.Permission
	for _, p := range form.Permission {
		permission = append(permission, &model.Permission{
			Model: orm.Model{ID: p},
		})
	}
	err := s.RoleService.Save(role, permission)
	if err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}

func (s *RoleCtrl) Page(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	ctx.Put("page", s.RoleService.Page(page))
	return ctx.ResultOK()
}
