package http

import (
	"github.com/labstack/echo"
)

const (
	SuffixJSON = ".json"
	SuffixHTML = ".html"
)

type Context struct {
	Ctx    echo.Context
	Suffix string
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

type HandlerFunc func(ctx *Context) error
