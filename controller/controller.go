package controller

import (
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/consts"
	"github.com/GoodHot/TinyCMS/controller/admin"
	"github.com/GoodHot/TinyCMS/controller/web"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Controller struct {
	Log               brain.Logger             `ioc:"auto"`
	AdminAuthCtrl     *admin.AdminAuthCtrl     `ioc:"auto"`
	AdminArticleCtrl  *admin.AdminArticleCtrl  `ioc:"auto"`
	AdminCategoryCtrl *admin.AdminCategoryCtrl `ioc:"auto"`
	AdminTagCtrl      *admin.AdminTagCtrl      `ioc:"auto"`
	AdminSkinCtrl     *admin.AdminSkinCtrl     `ioc:"auto"`
	AdminUploadCtrl   *admin.AdminUploadCtrl   `ioc:"auto"`
	AdminConfigCtrl   *admin.AdminConfigCtrl   `ioc:"auto"`
	IndexCtrl         *web.IndexCtrl           `ioc:"auto"`
	AdminService      *service.UserService     `ioc:"auto"`
	SkinService       *service.SkinService     `ioc:"auto"`
	ServerStatic      string                   `val:"${server.static}"`       // 静态文件夹
	ServerPort        int                      `val:"${server.port}"`         // 服务启动端口
	AdminPrefix       string                   `val:"${server.admin_prefix}"` // Admin接口访问路径
	WebPrefix         string                   `val:"${server.web_prefix}"`   // Web页面路径
}

func (s *Controller) Startup() error {
	// 初始化Echo
	e := echo.New()
	e.HideBanner = true
	// 模板引擎
	e.Renderer = s.SkinService
	// 静态文件
	e.Static(s.ServerStatic, s.ServerStatic)
	s.Log.Info("website static files is %s", s.ServerStatic)
	// 注册controller
	s.registerAdmin(e.Group(s.AdminPrefix))
	s.registerWeb(e.Group(s.WebPrefix))
	// 启动服务
	return e.Start(":" + strconv.Itoa(s.ServerPort))
}

func (s *Controller) registerWeb(group *echo.Group) {
	router := &RouterRegister{
		Group:        group,
		Prefix:       s.WebPrefix,
		BuildOption:  false,
		AdminService: s.AdminService,
		Log:          s.Log,
	}
	router.GET("/", s.IndexCtrl.Index)
	router.GET("/post/:shortId", s.IndexCtrl.Post)
	router.GET("/post/:shortId/:title", s.IndexCtrl.Post)
	router.GET("/tags", s.IndexCtrl.Tags)
	router.GET("/tag/:tag", s.IndexCtrl.Tag)
}

func (s *Controller) registerAdmin(group *echo.Group) {
	// 开发环境允许跨域
	group.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	router := &RouterRegister{
		Group:        group,
		Prefix:       s.AdminPrefix,
		BuildOption:  true,
		AdminService: s.AdminService,
		Log:          s.Log,
	}
	// 权限拦截
	group.Use(s.adminAuthInterceptor)

	// 配置controller
	// admin
	router.POST("/auth/login", s.AdminAuthCtrl.Login)
	// 临时使用
	router.GET("/auth/init", s.AdminAuthCtrl.InitAdmin)

	// user
	router.GET("/user/all", s.AdminAuthCtrl.All)

	// article
	router.POST("/article", s.AdminArticleCtrl.Publish)
	router.GET("/article/page_:page", s.AdminArticleCtrl.Page)
	router.GET("/article/:id", s.AdminArticleCtrl.Get)
	router.GET("/article/count", s.AdminArticleCtrl.Count)
	router.DELETE("/article", s.AdminArticleCtrl.Delete)

	// category
	router.POST("/category", s.AdminCategoryCtrl.Save)
	router.GET("/category/tree", s.AdminCategoryCtrl.Tree)
	router.GET("/category/:id", s.AdminCategoryCtrl.Get)
	router.GET("/category/page_:page", s.AdminCategoryCtrl.Page)
	router.GET("/category/name/:name", s.AdminCategoryCtrl.GetByName)
	router.GET("/category/path/:path", s.AdminCategoryCtrl.GetByPath)
	router.POST("/category/sort", s.AdminCategoryCtrl.Sort)

	// tag
	router.POST("/tag", s.AdminTagCtrl.Save)
	router.GET("/tag/:id", s.AdminTagCtrl.Get)
	router.GET("/tag/search/:prefix", s.AdminTagCtrl.Search)
	router.GET("/tag/hot", s.AdminTagCtrl.HotTag)
	router.GET("/tag/page_:page", s.AdminTagCtrl.Page)

	// skin
	router.GET("/skins", s.AdminSkinCtrl.List)
	router.PUT("/skin/switch/:id", s.AdminSkinCtrl.Switch)

	// upload
	router.POST("/upload", s.AdminUploadCtrl.Upload)
	router.POST("/upload/markdown", s.AdminUploadCtrl.MarkdownUpload)

	// config
	router.GET("/config/list/:group", s.AdminConfigCtrl.ListByGroup)
	router.PUT("/config", s.AdminConfigCtrl.Update)
}

func (s *Controller) adminAuthInterceptor(next echo.HandlerFunc) echo.HandlerFunc {
	whiteList := []string{"/auth/login", "/auth/init"}
	return func(c echo.Context) error {
		url := c.Request().URL.String()
		for _, v := range whiteList {
			if s.AdminPrefix+v == url {
				return next(c)
			}
		}
		token := c.Request().Header.Get(consts.HttpAccessTokenName())
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "please login")
		}
		_, err := s.AdminService.CheckToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "please login")
		}
		return next(c)
	}
}
