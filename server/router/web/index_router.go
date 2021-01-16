package web

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
)

type Index struct {
}

func (my *Index) Index(ctx *http.Context) *core.Err {
	ctx.Put("hello", "hello")
	return ctx.ResultOK("index")
}
