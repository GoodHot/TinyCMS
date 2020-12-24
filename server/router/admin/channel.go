package admin

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Channel struct {
	ChannelService *service.ChannelService `ioc:"auto"`
}

func (my *Channel) Save(ctx *http.Context) *core.Err {
	var channel trait.Channel
	if err := ctx.Bind(&channel); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	// TODO 验证channel合法性
	cn, _ := my.ChannelService.GetByPath(channel.Path)
	if cn != nil {
		return core.NewErr(core.Err_Channel_Path_Exists)
	}
	if err := my.ChannelService.Save(&channel); err != nil {
		return err
	}
	return ctx.ResultOK("ok")
}
