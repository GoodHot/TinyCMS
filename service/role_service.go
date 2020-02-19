package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type RoleService struct {
	ORM          *orm.ORM      `ioc:"auto"`
	AdminService *AdminService `ioc:"auto"`
}

func (s *RoleService) Init() {

	role := &model.Role{
		Code:     "super",
		RoleName: "超级管理员",
		IsSuper:  true,
	}
	s.ORM.DB.Save(role)
	s.ORM.DB.Save(&model.Admin{
		Nickname: "admin",
		Username: "admin",
		Password: s.AdminService.encrptyPwd("admin"),
		Resume:   "super admin",
		RoleID:   role.ID,
	})

	content := &model.Permission{
		Code:           "content",
		PermissionName: "内容管理",
		PID:            0,
	}
	s.ORM.DB.Save(content)
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_publish",
		PermissionName: "发布文章",
		PID:            content.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_edit",
		PermissionName: "编辑文章",
		PID:            content.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_delete",
		PermissionName: "删除文章",
		PID:            content.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_list",
		PermissionName: "查询文章列表",
		PID:            content.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_channel_edit",
		PermissionName: "编辑分类",
		PID:            content.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_channel_delete",
		PermissionName: "删除分类",
		PID:            content.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "content_channel_save",
		PermissionName: "增加分类",
		PID:            content.ID,
	})
}
