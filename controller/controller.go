package controller

import (
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/render"
	"github.com/GoodHot/TinyCMS/controller/admin"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	Log                brain.Logger             `ioc:"auto"`
	AdminAuthCtrl      *admin.AdminAuthCtrl     `ioc:"auto"`
	AdminArticleCtrl   *admin.AdminArticleCtrl  `ioc:"auto"`
	AdminCategoryCtrl  *admin.AdminCategoryCtrl `ioc:"auto"`
	AdminTagCtrl       *admin.AdminTagCtrl      `ioc:"auto"`
	AdminService       *service.AdminService    `ioc:"auto"`
	ServerStatic       string                   `val:"${server.static}"`            // 静态文件夹
	ServerPort         int                      `val:"${server.port}"`              // 服务启动端口
	ServerTemplateDir  string                   `val:"${server.skin_template_dir}"` // 模板文件夹
	ServerHTMLCompress bool                     `val:"${server.html_compress}"`     // HTML代码是否压缩
	AdminPrefix        string                   `val:"${server.admin_prefix}"`      // Admin接口访问路径
}

func (s *Controller) Startup() error {
	// 初始化Echo
	e := echo.New()
	e.HideBanner = true
	// 模板引擎
	e.Renderer = &render.TinyRender{
		TemplateDir:  s.ServerTemplateDir,
		HTMLCompress: s.ServerHTMLCompress,
	}
	// 静态文件
	e.Static(s.ServerStatic, s.ServerStatic)
	s.Log.Info("website static files is %s", s.ServerStatic)
	// 注册controller
	s.registerAdmin(e.Group(s.AdminPrefix))
	// 启动服务
	return e.Start(":" + strconv.Itoa(s.ServerPort))
}

func (s *Controller) registerAdmin(group *echo.Group) {
	router := &RouterRegister{
		Group:        group,
		Prefix:       s.AdminPrefix,
		BuildOption:  false,
		AdminService: s.AdminService,
		Log:          s.Log,
	}
	// 权限拦截
	group.Use(s.adminAuthInterceptor)

	// 配置controller
	// admin
	router.POST("/auth/login", s.AdminAuthCtrl.Login)

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
	router.POST("/category/sort", s.AdminCategoryCtrl.Sort)

	// tag
	router.POST("/tag", s.AdminTagCtrl.Save)
	router.GET("/tag/:id", s.AdminTagCtrl.Get)
	router.GET("/tag/search/:prefix", s.AdminTagCtrl.Search)
	router.GET("/tag/hot", s.AdminTagCtrl.HotTag)
	router.GET("/tag/page_:page", s.AdminTagCtrl.Page)
}

func (s *Controller) adminAuthInterceptor(next echo.HandlerFunc) echo.HandlerFunc {
	whiteList := []string{"/auth/login"}
	return func(c echo.Context) error {
		url := c.Request().URL.String()
		for _, v := range whiteList {
			if s.AdminPrefix+v == url {
				return next(c)
			}
		}
		token := c.Request().Header.Get("ACCESS-TOKEN")
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
