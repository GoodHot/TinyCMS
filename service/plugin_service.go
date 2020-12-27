package service

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type ParamOperator struct {
	pid int
	orm *orm.ORMFactory
}

func (p *ParamOperator) Set(key string, value string, canEdit bool) {
}

type IPlugin interface {
	Install(param *ParamOperator)
}

type PluginService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (ps *PluginService) Mount(server string, internal bool) *core.Err {
	// TODO 发起http请求
	// TODO 获取保存参数
	plugin := &trait.Plugin{
		Name:        "LocalUpload",
		Version:     "1.0.0",
		Description: "local upload file",
		Type:        trait.PluginTypeUpload,
		Server:      server,
		Internal:    internal,
	}
	dbPlugin, _ := ps.ORM.Plugin.GetByInfo(plugin.Name, plugin.Version, plugin.Type)
	if dbPlugin != nil {
		return core.NewErr(core.Err_Plugin_Exists)
	}
	ps.ORM.Plugin.Save(plugin)
	// 保存参数设置
	return nil
}

func (ps *PluginService) MountInternal(plugin *trait.Plugin, iplugin IPlugin) *core.Err {
	// TODO 验证插件合法性
	plugin.Internal = true
	dbPlugin, _ := ps.ORM.Plugin.GetByInfo(plugin.Name, plugin.Version, plugin.Type)
	if dbPlugin != nil {
		return core.NewErr(core.Err_Plugin_Exists)
	}
	ps.ORM.Plugin.Save(plugin)
	param := &ParamOperator{orm: ps.ORM, pid: plugin.ID}
	iplugin.Install(param)
	return nil
}
