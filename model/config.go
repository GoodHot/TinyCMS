package model

import (
	"github.com/GoodHot/TinyCMS/orm"
	"strconv"
)

type OptType string
type ConfigGroup string

const (
	OptTypeText        OptType = "text"
	OptTypeTextArea    OptType = "textarea"
	OptTypeSelect      OptType = "select" // 格式为 key:value,key:value
	OptTypeUploadImage OptType = "uploadImage"
	OptTypeTags        OptType = "tags"

	DictGroupSystem ConfigGroup = "system" // 系统配置
	DictGroupSEO    ConfigGroup = "seo"    // 用户配置
)

type Config struct {
	orm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Key         string      `json:"key"`
	Value       string      `json:"value"`
	OptType     OptType     `json:"opt_type"`     // 操作类型
	Options     string      `json:"options"`      // 操作选项
	Visible     bool        `json:"visible"`      // 是否可见
	ConfigGroup ConfigGroup `json:"config_group"` // 配置分组
}

func (s *Config) BoolVal() bool {
	val, _ := strconv.ParseBool(s.Value)
	return val
}

func (s *Config) IntVal() int {
	val, _ := strconv.Atoi(s.Value)
	return val
}
