package web

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
)

type Post struct {
	PostService *service.PostService `ioc:"auto"`
}

func (my *Post) Post(ctx *http.Context) *core.Err {
	fmt.Println(ctx.Param("id"))
	return ctx.ResultOK("post")
}
