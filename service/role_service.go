package service

import (
	"errors"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/jinzhu/gorm"
)

type RoleService struct {
	ORM          *orm.ORM      `ioc:"auto"`
	AdminService *AdminService `ioc:"auto"`
	PageSize     int           `val:"${website.page_size}"`
	permissions  map[uint]*model.Permission
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

func (s *RoleService) Save(role model.Role, permissions []*model.Permission) error {
	exists := &model.Role{}
	s.ORM.DB.Where("code = ? and id != ?", role.Code, role.ID).Find(exists)
	if exists.ID != 0 {
		return errors.New(strs.Fmt("角色编码：%s 已经存在", role.Code))
	}
	if role.IsSuper {
		return s.ORM.DB.Save(&role).Error
	}
	return s.ORM.Tx(func(db *gorm.DB) error {
		if role.ID > 0 {
			db.Unscoped().Where("role_id = ?", role.ID).Delete(&model.RelRolePermission{})
		}
		if err := db.Save(&role).Error; err != nil {
			return err
		}
		for _, per := range permissions {
			rel := model.RelRolePermission{
				RoleID:     role.ID,
				Permission: per.ID,
			}
			if err := db.Save(&rel).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *RoleService) Page(page int) *orm.Page {
	var roles []*model.Role
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &roles,
		OrderBy:  "is_super desc, id desc",
	})
	if roles != nil && len(roles) > 0 {
		s.getRelsPermission(&roles)
	}
	return result
}

func (s *RoleService) getRelsPermission(roles *[]*model.Role) {
	for _, role := range *roles {
		s.getRelPermission(role)
	}
}

func (s *RoleService) getRelPermission(role *model.Role) {
	var rels []*model.RelRolePermission
	s.ORM.DB.Where("role_id = ?", role.ID).Find(&rels)
	if rels == nil || len(rels) == 0 {
		return
	}
	permissions := role.Permissions
	if permissions == nil {
		role.Permissions = make(map[string][]*model.Permission)
	}
	maps := s.permissionMap()
	for _, rel := range rels {
		value := maps[rel.Permission]
		if value.Parent != nil {
			parentMaps := role.Permissions[value.Parent.PermissionName]
			if parentMaps == nil {
				parentMaps = []*model.Permission{}
			}
			parentMaps = append(parentMaps, value)
			role.Permissions[value.Parent.PermissionName] = parentMaps
		}
	}
}

func (s *RoleService) permissionMap() map[uint]*model.Permission {
	if s.permissions != nil {
		return s.permissions
	}
	var pers []*model.Permission
	s.ORM.DB.Find(&pers)
	maps := make(map[uint]*model.Permission)
	for _, parent := range pers {
		if parent.PID == 0 {
			maps[parent.ID] = parent
		}
	}
	for _, child := range pers {
		if child.PID == 0 {
			continue
		}
		maps[child.ID] = child
		child.Parent = maps[child.PID]
	}
	s.permissions = maps
	return maps
}

func (s *RoleService) Get(id int) (*model.Role, error) {
	var role model.Role
	s.ORM.DB.Where("id = ?", id).First(&role)
	if &role == nil {
		return nil, errors.New("角色不存在")
	}
	s.getRelPermission(&role)
	return &role, nil
}

func (s *RoleService) All() []*model.Role {
	var roles []*model.Role
	s.ORM.DB.Order("id asc").Find(&roles)
	return roles
}
