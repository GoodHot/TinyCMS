package http

import "github.com/gin-gonic/gin"

type Context struct {
	Ctx *gin.Context
}

// 获取ip地址
func (*Context) IPAddr() string {
	return ""
}

type HandlerFunc func(ctx *Context) error
