package router

import (
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/server/router/admin"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/server/router/web"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	http2 "net/http"
	"strings"
)

type Router struct {
	JWTSecret       string                   `val:"${props.secret.jwt}"`
	RouterAdmin     string                   `val:"${props.router.admin}"`
	RouterWeb       string                   `val:"${props.router.web}"`
	RouterIndex     string                   `val:"${props.router.index}"`
	TemplateService *service.TemplateService `ioc:"auto"`
	AdminMember     *admin.Member            `ioc:"auto"`
	AdminChannel    *admin.Channel           `ioc:"auto"`
	AdminPlugin     *admin.Plugin            `ioc:"auto"`
	AdminDict       *admin.Dict              `ioc:"auto"`
	AdminPost       *admin.Post              `ioc:"auto"`
	AdminUpload     *admin.Upload            `ioc:"auto"`
	WebIndex        *web.Index               `ioc:"auto"`
	WebChannel      *web.Channel             `ioc:"auto"`
	WebPost         *web.Post                `ioc:"auto"`
}

func (router *Router) Register(echo *echo.Echo, group *echo.Group) {
	echo.Renderer = router.TemplateService
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
				Internal: nil,
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
	register.POST("/auth/signin", router.AdminMember.SignIn)
	// 获取当前用户
	register.GET("/auth/info", router.AdminMember.Info)
	// 保存频道
	register.POST("/channel", router.AdminChannel.Save)
	// 获取所有频道列表
	register.GET("/channels", router.AdminChannel.All)
	// 修改频道排序
	register.PUT("/channels/sort", router.AdminChannel.Sort)
	// 挂载插件
	register.POST("/plugin/mount", router.AdminPlugin.Mount)
	// 根据类型获取插件列表
	register.GET("/plugin/:type/all", router.AdminPlugin.GetByType)
	// 修改插件参数
	register.PUT("/plugin/params", router.AdminPlugin.UpdateParam)
	// 更新字典
	register.PUT("/dict", router.AdminDict.Update)
	// 查询字典
	register.GET("/dict/:key/info", router.AdminDict.Get)
	// 保存文章
	register.POST("/post", router.AdminPost.Save)
	// 文章列表
	register.GET("/post/:page/:size/page", router.AdminPost.Page)
	// 文件上传
	register.POST("/upload", router.AdminUpload.Update)
	// 获取管理员列表
	register.GET("/staff/:page/:size/page", router.AdminMember.StaffList)
}

func (router *Router) registerWeb(group *echo.Group) {
	register := http.RouterRegister{Group: group, Index: router.RouterIndex}
	// 首页
	register.GET("/index", router.WebIndex.Index)
	// channel列表
	register.GET("/channel", router.WebChannel.Channel)
	// channel列表详情
	register.GET("/channel/:id", router.WebChannel.Channel)
	// post
	register.GET("/post/:id", router.WebPost.Post)
}
