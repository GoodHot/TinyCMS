package model

import "github.com/GoodHot/TinyCMS/orm"

type Role struct {
	orm.Model
	Code     string `gorm:"type:varchar(50);unique_index" json:"code"`
	RoleName string `json:"role_name"`
	IsSuper  bool   `json:"is_super"`
}

type RelRolePermission struct {
	RoleID     uint
	Permission uint
}

type Permission struct {
	orm.Model
	Code           string        `json:"code"`            // 唯一识别码
	PermissionName string        `json:"permission_name"` // 权限名称
	PID            uint          // 上级ID
	Child          []*Permission `gorm:"-" json:"child"`
}
