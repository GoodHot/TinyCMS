package service

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type PluginService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (ps *PluginService) Mount(server string) *core.Err {
	// 发起http请求
	// 获取保存参数
	plugin := &trait.Plugin{
		Name:        "LocalUpload",
		Version:     "1.0.0",
		Description: "local upload file",
		Type:        trait.PluginTypeUpload,
		Server:      server,
		Internal:    false,
	}
	dbPlugin, _ := ps.ORM.Plugin.GetByInfo(plugin.Name, plugin.Version, plugin.Type)
	if dbPlugin != nil {
		return core.NewErr(core.Err_Plugin_Exists)
	}
	ps.ORM.Plugin.Save(plugin)
	// 保存参数设置
	return nil
}
