package http

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RouterRegister struct {
	Group  *echo.Group
	Prefix string
	Index  string
}

func (reg *RouterRegister) Middleware(middleware echo.MiddlewareFunc) {
	reg.Group.Use(middleware)
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

func (reg *RouterRegister) errorHandler(ctx *Context, err *core.Err) error {
	if err == nil {
		return nil
	}
	result := NewRespResult(int(err.ErrType.Code), err.Error(), nil)
	return ctx.Ctx.JSON(err.ErrType.RespStatus, result)
}

func (reg *RouterRegister) register(method string, relativePath, suffix string, handlerFunc HandlerFunc) {
	fmt.Printf("register:%v%v\n", relativePath, suffix)
	relativePath = relativePath + suffix
	if method == http.MethodGet {
		if reg.Index+SuffixHTML == relativePath && suffix == SuffixHTML {
			fmt.Printf("register index \\")
			reg.Group.GET("/", func(context echo.Context) error {
				ctx := &Context{
					Ctx:    context,
					Suffix: SuffixHTML,
				}
				return reg.errorHandler(ctx, handlerFunc(ctx))
			})
		}
		reg.Group.GET(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else if method == http.MethodHead {
		reg.Group.HEAD(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else if method == http.MethodPost {
		reg.Group.POST(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else if method == http.MethodPut {
		reg.Group.PUT(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else if method == http.MethodPatch {
		reg.Group.PATCH(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else if method == http.MethodDelete {
		reg.Group.DELETE(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else if method == http.MethodOptions {
		reg.Group.OPTIONS(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	} else {
		reg.Group.Any(relativePath, func(context echo.Context) error {
			ctx := &Context{
				Ctx:    context,
				Suffix: suffix,
			}
			return reg.errorHandler(ctx, handlerFunc(ctx))
		})
	}
}
