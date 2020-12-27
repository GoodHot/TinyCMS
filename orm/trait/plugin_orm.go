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

type PluginParam struct {
	BaseORM
	PID         int
	Key         string
	Value       string
	Description string
	CanEdit     bool
	Visible     bool
}

type PluginORM interface {
	GetByInfo(name string, version string, pluginType PluginType) (*Plugin, *core.Err)
	Save(plugin *Plugin) *core.Err
}
