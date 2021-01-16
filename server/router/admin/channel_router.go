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

func (my *Channel) All(ctx *http.Context) *core.Err {
	cs, err := my.ChannelService.GetAll()
	if err != nil {
		return err
	}
	ctx.Put("channels", cs)
	return ctx.ResultOK("ok")
}

type ChannelSort struct {
	Sortable []struct {
		ID   int `json:"id"`
		Sort int `json:"sort"`
	} `json:"sortable"`
}

func (my *Channel) Sort(ctx *http.Context) *core.Err {
	var cs ChannelSort
	err := ctx.Bind(&cs)
	if err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	var data []trait.Channel
	for _, sort := range cs.Sortable {
		data = append(data, trait.Channel{
			BaseORM: trait.BaseORM{
				ID: sort.ID,
			},
			Sort: sort.Sort,
		})
	}
	my.ChannelService.UpdateSort(data)
	return ctx.ResultOK("ok")
}
