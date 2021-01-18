package web

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Post struct {
	PostService *service.PostService `ioc:"auto"`
}

func (my *Post) Post(ctx *http.Context) *core.Err {
	postID := ctx.ParamInt("id")
	post, err := my.PostService.Get(postID)
	if err != nil {
		return err
	}
	ctx.Put("post", post)
	return ctx.ResultOK("post")
}
