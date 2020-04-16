package model

import (
	"github.com/GoodHot/TinyCMS/orm"
	"strconv"
)

type OptType string
type ConfigType string

const (
	OptTypeText        OptType = "text"
	OptTypeTextArea    OptType = "textarea"
	OptTypeSelect      OptType = "select" // 格式为 key:value,key:value
	OptTypeUploadImage OptType = "uploadImage"

	DictTypeSystem ConfigType = "system" // 系统配置
	DictTypeUser   ConfigType = "user"   // 用户配置
)

type Config struct {
	orm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Key         string     `json:"key"`
	Value       string     `json:"value"`
	OptType     OptType    `json:"opt_type"`    // 操作类型
	Options     string     `json:"options"`     // 操作选项
	Visible     bool       `json:"visible"`     // 是否可见
	ConfigType  ConfigType `json:"config_type"` // 配置类型
}

func (s *Config) BoolVal() bool {
	val, _ := strconv.ParseBool(s.Value)
	return val
}

func (s *Config) IntVal() int {
	val, _ := strconv.Atoi(s.Value)
	return val
}
