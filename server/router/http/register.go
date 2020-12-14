package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterRegister struct {
	Group *gin.RouterGroup
}

func (reg *RouterRegister) POST(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodPost, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodPost, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) GET(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodGet, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodGet, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) DELETE(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodDelete, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodDelete, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) PATCH(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodPatch, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodPatch, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) PUT(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodPut, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodPut, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) OPTIONS(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodOptions, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodOptions, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) HEAD(relativePath string, handlerFunc HandlerFunc) {
	reg.register(http.MethodHead, relativePath, SuffixJSON, handlerFunc)
	reg.register(http.MethodHead, relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) Any(relativePath string, handlerFunc HandlerFunc) {
	reg.register("", relativePath, SuffixJSON, handlerFunc)
	reg.register("", relativePath, SuffixHTML, handlerFunc)
}

func (reg *RouterRegister) register(method string, relativePath, suffix string, handlerFunc HandlerFunc) {
	relativePath = relativePath + suffix
	if method == http.MethodGet {
		reg.Group.POST(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	} else if method == http.MethodHead {
		reg.Group.HEAD(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			handlerFunc(ctx)
		})
	} else if method == http.MethodPost {
		reg.Group.POST(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	} else if method == http.MethodPut {
		reg.Group.PUT(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	} else if method == http.MethodPatch {
		reg.Group.PATCH(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	} else if method == http.MethodDelete {
		reg.Group.DELETE(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	} else if method == http.MethodOptions {
		reg.Group.OPTIONS(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	} else {
		reg.Group.Any(relativePath, func(context *gin.Context) {
			ctx := &Context{
				Ctx: context,
			}
			handlerFunc(ctx)
		})
	}
}
