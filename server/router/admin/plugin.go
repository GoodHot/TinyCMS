package admin

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Plugin struct {
	PluginService *service.PluginService `ioc:"auto"`
}

type pluginMount struct {
	PluginServer string `json:"plugin_server"`
}

func (my *Plugin) Mount(ctx *http.Context) *core.Err {
	var mount pluginMount
	if err := ctx.Bind(&mount); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	return my.PluginService.Mount(mount.PluginServer, false)
}
