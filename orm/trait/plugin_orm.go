package trait

import "github.com/GoodHot/TinyCMS/core"

type PluginType string

const (
	PluginTypeUpload PluginType = "upload"
)

type Plugin struct {
	BaseORM
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Type        PluginType `json:"type"`
	Server      string     `json:"server"`   // 插件服务器
	Internal    bool       `json:"internal"` // 是否为内部插件 true为内部，false为外部
}

type PluginParamType string

const (
	PluginParamType_Input    PluginParamType = "input"
	PluginParamType_Textarea PluginParamType = "textarea"
	PluginParamType_Select   PluginParamType = "select"
)

type PluginParam struct {
	BaseORM
	PID         int             `json:"pid"`
	Key         string          `json:"key"`
	Value       string          `json:"value"`
	Description string          `json:"description"`
	CanEdit     bool            `json:"can_edit"`
	Visible     bool            `json:"visible"`
	ValueType   PluginParamType `json:"value_type"`
}

type PluginORM interface {
	GetByInfo(name string, version string, pluginType PluginType) (*Plugin, *core.Err)
	Save(plugin *Plugin) *core.Err
	SaveParam(param *PluginParam) *core.Err
	GetByType(pluginType string) (*[]Plugin, *core.Err)
	GetParams(id int) *[]PluginParam
}
