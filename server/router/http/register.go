package http

import (
	"github.com/gin-gonic/gin"
)

type RouterRegister struct {
	Group *gin.RouterGroup
}

func (reg *RouterRegister) POST(relativePath string, handlerFunc HandlerFunc) {
}

func (reg *RouterRegister) GET(relativePath string, handlerFunc HandlerFunc) {
	reg.Group.GET(relativePath, func(context *gin.Context) {
		ctx := &Context{
			Ctx: context,
		}
		handlerFunc(ctx)
	})
}

func (reg *RouterRegister) DELETE(relativePath string, handlerFunc HandlerFunc) {
}

func (reg *RouterRegister) PATCH(relativePath string, handlerFunc HandlerFunc) {
}

func (reg *RouterRegister) PUT(relativePath string, handlerFunc HandlerFunc) {
}

func (reg *RouterRegister) OPTIONS(relativePath string, handlerFunc HandlerFunc) {
}

func (reg *RouterRegister) HEAD(relativePath string, handlerFunc HandlerFunc) {
}

func (reg *RouterRegister) Any(relativePath string, handlerFunc HandlerFunc) {
}
