package router

import (
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/server/router/admin"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	http2 "net/http"
	"strings"
)

type Router struct {
	JWTSecret    string         `val:"${props.secret.jwt}"`
	RouterAdmin  string         `val:"${props.router.admin}"`
	RouterWeb    string         `val:"${props.router.web}"`
	AdminAdmin   *admin.Admin   `ioc:"auto"`
	AdminChannel *admin.Channel `ioc:"auto"`
}

func (router *Router) Register(group *echo.Group) {
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
			url := context.Request().URL
			uri := url.Path[:strings.LastIndex(url.Path, ".")]
			return uri == prefix+"/auth/signin"
		},
		ErrorHandler: func(err error) error {
			return &echo.HTTPError{
				Code:     http2.StatusUnauthorized,
				Message:  "please login",
				Internal: err,
			}
		},
		ContextKey:    config.JWTContextKey,
		SigningMethod: middleware.AlgorithmHS256,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningKey:    []byte(router.JWTSecret),
		Claims:        jwt.MapClaims{},
		AuthScheme:    "TinyCMS",
	}))
	// 登录
	register.POST("/auth/signin", router.AdminAdmin.Signin)
	// 获取当前用户
	register.GET("/auth/info", router.AdminAdmin.Info)
	// 保存频道
	register.POST("/channel", router.AdminChannel.Save)
	// 获取所有频道列表
	register.GET("/channels", router.AdminChannel.All)
	// 修改频道排序
	register.PUT("/channels/sort", router.AdminChannel.Sort)
	// 获取当前用户
	//register.POST("/user", router.AdminIndex.Index)
	//// 获取文章列表 :status = [published(default), drafts, scheduled]
	//register.GET("/posts/:status", router.AdminIndex.List)
	//// 获取频道
	//register.GET("/channels", router.AdminIndex.List)
	//// 获取标签
	//register.GET("/tags", router.AdminIndex.List)
	//// 获取页面
	//register.GET("/pages", router.AdminIndex.List)
	//// 站点设置
	//register.GET("/general", router.AdminIndex.List)
	//// 修改站点设置
	//register.PUT("/general/:key", router.AdminIndex.List)
	//// 页面设计
	//register.GET("/design", router.AdminIndex.List)
	//// 代码块注入
	//register.GET("/code-injection", router.AdminIndex.List)

}

func (router *Router) registerWeb(group *echo.Group) {
	//register := http.RouterRegister{Group: group}
	//// 首页
	//register.GET("/", router.AdminIndex.List)
	//// 查询文章列表
	//register.GET("/posts/:page", router.AdminIndex.List)
	//// 查询文章详情
	//register.GET("/post/:name", router.AdminIndex.List)
	//// 查询所有页面
	//register.GET("/pages", router.AdminIndex.List)
	//// 查询页面详情
	//register.GET("/page/:name", router.AdminIndex.List)
	//// 查询所有标签
	//register.GET("/tags", router.AdminIndex.List)
	//// 查询某个标签下所有文章
	//register.GET("/tag/:name", router.AdminIndex.List)

}
