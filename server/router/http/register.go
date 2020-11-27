package http

import (
	"github.com/gin-gonic/gin"
)

type RouterRegister struct {
	Group *gin.RouterGroup
}

func (reg *RouterRegister) POST(relativePath string) {
}

func (reg *RouterRegister) GET(relativePath string, handlerFunc HandlerFunc) {
	reg.Group.GET(relativePath, func(context *gin.Context) {
		ctx := &Context{
			Ctx: context,
		}
		handlerFunc(ctx)
	})
}

func (reg *RouterRegister) DELETE(relativePath string) {
}

func (reg *RouterRegister) PATCH(relativePath string) {
}

func (reg *RouterRegister) PUT(relativePath string) {
}

func (reg *RouterRegister) OPTIONS(relativePath string) {
}

func (reg *RouterRegister) HEAD(relativePath string) {
}

func (reg *RouterRegister) Any(relativePath string) {
}
