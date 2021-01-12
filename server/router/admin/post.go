package admin

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Post struct {
	PostService *service.PostService `ioc:"auto"`
}

func (my *Post) Save(ctx *http.Context) *core.Err {
	var form service.SavePostForm
	if err := ctx.Bind(&form); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	my.PostService.Save(&form, ctx.CurrMember().ID)
	return ctx.ResultOK("ok")
}

func (my *Post) Page(ctx *http.Context) *core.Err {
	page := ctx.ParamInt("page")
	size := ctx.ParamInt("size")
	lastID := ctx.QueryParamInt("lastID")
	channelID := ctx.QueryParamInt("channelID")
	status := ctx.QueryParamInt("status")
	query := &trait.Query{
		Page:   page,
		Size:   size,
		LastID: lastID,
		Param: map[string]interface{}{
			"channel_id": channelID,
			"status":     status,
		},
		Order: map[string]string{
			"id": "desc",
		},
	}
	result, err := my.PostService.Page(query)
	if err != nil {
		return core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	ctx.Put("page", result)
	return ctx.ResultOK("ok")
}
