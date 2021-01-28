package router

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	http2 "net/http"
)

type Router struct {
	JWTSecret   string `val:"${props.secret.jwt}"`
	RouterAdmin string `val:"${props.router.admin}"`
	RouterWeb   string `val:"${props.router.web}"`
	RouterIndex string `val:"${props.router.index}"`
}

func (router *Router) Register(echo *echo.Echo, group *echo.Group) {
	// 注册admin路由
	adminRouter := group.Group(router.RouterAdmin)
	router.registerAdmin(adminRouter, router.RouterAdmin)
	// 注册web路由
	webRouter := group.Group(router.RouterWeb)
	router.registerWeb(webRouter)
}

func (router *Router) registerAdmin(group *echo.Group, prefix string) {
	register := http.RouterRegister{Group: group, Prefix: prefix}
	register.Middleware(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(context echo.Context) bool {
			uri := context.Request().RequestURI
			fmt.Println(uri)
			return uri == prefix+"/auth/signin/json"
		},
		ErrorHandler: func(err error) error {
			return &echo.HTTPError{
				Code:     http2.StatusUnauthorized,
				Message:  "please login",
				Internal: nil,
			}
		},
		ContextKey:    "config.JWTContextKey",
		SigningMethod: middleware.AlgorithmHS256,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningKey:    []byte(router.JWTSecret),
		Claims:        jwt.MapClaims{},
		AuthScheme:    "TinyCMS",
	}))
}

func (router *Router) registerWeb(group *echo.Group) {
	//register := http.RouterRegister{Group: group, Index: router.RouterIndex}
}
