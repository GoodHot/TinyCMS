package web

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
)

type Channel struct {
}

func (my *Channel) Channel(ctx *http.Context) *core.Err {
	return ctx.ResultOK("channel")
}

