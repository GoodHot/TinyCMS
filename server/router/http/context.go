package http

import (
	"github.com/labstack/echo"
	"net/http"
)

const (
	SuffixJSON = ".json"
	SuffixHTML = ".html"
)

type Context struct {
	Ctx    echo.Context
	Suffix string
	data   map[string]interface{}
}

func (ctx *Context) Put(key string, val interface{}) {
	if ctx.data == nil {
		ctx.data = make(map[string]interface{})
	}
	ctx.data[key] = val
}

// 获取ip地址
func (*Context) IPAddr() string {
	return ""
}

func (ctx *Context) HTML(html string) error {
	ctx.Ctx.String(200, html)
	return nil
}

func (ctx *Context) Bind(i interface{}) error {
	return ctx.Ctx.Bind(i)
}

func (ctx *Context) JSON() error {
	return ctx.Ctx.JSON(http.StatusOK, ctx.data)
}

type HandlerFunc func(ctx *Context) error
