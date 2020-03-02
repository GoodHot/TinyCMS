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

	article := &model.Permission{
		Code:           "article",
		PermissionName: "文章管理",
		PID:            0,
	}
	s.ORM.DB.Save(article)
	s.ORM.DB.Save(&model.Permission{
		Code:           "article_publish",
		PermissionName: "发布文章",
		PID:            article.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "article_edit",
		PermissionName: "编辑文章",
		PID:            article.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "article_delete",
		PermissionName: "删除文章",
		PID:            article.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "article_list",
		PermissionName: "查询文章列表",
		PID:            article.ID,
	})

	category := &model.Permission{
		Code:           "category",
		PermissionName: "分类管理",
		PID:            0,
	}
	s.ORM.DB.Save(category)
	s.ORM.DB.Save(&model.Permission{
		Code:           "category_edit",
		PermissionName: "编辑分类",
		PID:            category.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "category_delete",
		PermissionName: "删除分类",
		PID:            category.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "category_save",
		PermissionName: "增加分类",
		PID:            category.ID,
	})
	s.ORM.DB.Save(&model.Permission{
		Code:           "category_query",
		PermissionName: "查询分类",
		PID:            category.ID,
	})
}

func (s *RoleService) AllPermission() []*model.Permission {
	var parent []*model.Permission
	s.ORM.DB.Where("p_id = ?", 0).Find(&parent)
	if parent != nil && len(parent) > 0 {
		for _, p := range parent {
			var child []*model.Permission
			s.ORM.DB.Where("p_id = ?", p.ID).Find(&child)
			p.Child = child
		}
	}
	return parent
}
