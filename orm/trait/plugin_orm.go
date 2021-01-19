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

type FormType string

const (
	FormType_Input    FormType = "input"
	FormType_Switch   FormType = "switch"
	FormType_Number   FormType = "number"
	FormType_Textarea FormType = "textarea"
	FormType_Select   FormType = "select"
	FormType_Tags     FormType = "tags"
)

type PluginParam struct {
	BaseORM
	PID         int      `json:"pid"`
	Key         string   `json:"key"`
	Value       string   `json:"value"`
	Description string   `json:"description"`
	CanEdit     bool     `json:"can_edit"`
	Visible     bool     `json:"visible"`
	ValueType   FormType `json:"value_type"`
}

type PluginORM interface {
	GetByInfo(name string, version string, pluginType PluginType) (*Plugin, *core.Err)
	Save(plugin *Plugin) *core.Err
	SaveParam(param *PluginParam) *core.Err
	GetByType(pluginType string) (*[]Plugin, *core.Err)
	GetParams(id int) *[]PluginParam
	UpdateParam(param PluginParam) *core.Err
}
