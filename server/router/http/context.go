package http

import "github.com/gin-gonic/gin"

const (
	SuffixJSON = ".json"
	SuffixHTML = ".html"
)

type Context struct {
	Ctx    *gin.Context
	Suffix string
}

// 获取ip地址
func (*Context) IPAddr() string {
	return ""
}

func (ctx *Context) ShouldBind(i interface{}) error {
	return ctx.Ctx.ShouldBindUri(i)
}

func (ctx *Context) HTML(html string) error {
	ctx.Ctx.String(200, html)
	return nil
}

func (ctx *Context) BindJSON(i interface{}) error {
	return ctx.Ctx.BindJSON(i)
}

type HandlerFunc func(ctx *Context) error
