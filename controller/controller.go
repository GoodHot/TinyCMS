package controller

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/controller/admin"
	"github.com/GoodHot/TinyCMS/controller/web"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	"net/http"
	"strconv"
)

type Controller struct {
	AdminCategoryCtrl *admin.AdminCategoryCtrl `ioc:"auto"`
	AdminTagCtrl      *admin.AdminTagCtrl      `ioc:"auto"`
	AdminAuthCtrl     *admin.AdminAuthCtrl     `ioc:"auto"`
	AdminArticleCtrl  *admin.AdminArticleCtrl  `ioc:"auto"`
	AdminUploadCtrl   *admin.AdminUploadCtrl   `ioc:"auto"`
	AdminDictCtrl     *admin.AdminDictCtrl     `ioc:"auto"`
	RoleAuthCtrl      *admin.RoleCtrl          `ioc:"auto"`
	SkinCtrl          *web.SkinCtrl            `ioc:"auto"`
	IndexCtrl         *web.IndexCtrl           `ioc:"auto"`
	AdminService      *service.AdminService    `ioc:"auto"`
	Port              int                      `val:"${server.port}"`
	AdminPrefix       string                   `val:"${server.admin_prefix}"`
	WebPrefix         string                   `val:"${server.web_prefix}"`
	APIPrefix         string                   `val:"${server.api_prefix}"`
	Static            string                   `val:"${server.static}"`
	HTMLCompress      bool                     `val:"${server.html_compress}"`
	AdminTemplateDir  string                   `val:"${server.admin_template_dir}"`
}

func (c *Controller) StartUp() {
	e := echo.New()
	e.HideBanner = true
	e.Renderer = c
	e.Static(c.Static, c.Static)

	c.registerAdmin(e.Group(c.AdminPrefix), c.AdminPrefix)
	c.registerWeb(e, c.WebPrefix)
	c.registerAPI(e.Group(c.APIPrefix), c.APIPrefix)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(c.Port)))
	//if err != nil {
	//	panic(err)
	//}
}

func (s *Controller) registerAdmin(group *echo.Group, prefix string) {
	group.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	whiteList := []string{"/auth/login"}

	group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			url := c.Request().URL.String()
			for _, v := range whiteList {
				if prefix+v == url {
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
	})

	router := &RouterRegister{
		group:       group,
		prefix:      prefix,
		buildOption: true,
	}
	// auth
	router.POST("/auth/login", s.AdminAuthCtrl.Login, s.AdminService)
	router.Any("/auth/info", s.AdminAuthCtrl.Info, s.AdminService)

	// admin
	router.GET("/admin/all", s.AdminAuthCtrl.All, s.AdminService)
	router.POST("/admin", s.AdminAuthCtrl.Save, s.AdminService)
	router.GET("/admin/:id", s.AdminAuthCtrl.Get, s.AdminService)

	// channel
	router.POST("/category", s.AdminCategoryCtrl.Save, s.AdminService)
	router.GET("/category/tree", s.AdminCategoryCtrl.Tree, s.AdminService)
	//router.GET("/channel/page_:page", s.AdminChannelCtrl.Page)

	// tag
	router.GET("/tag/search/:prefix", s.AdminTagCtrl.Search, s.AdminService)
	router.GET("/tag/hot", s.AdminTagCtrl.HotTag, s.AdminService)

	// upload
	router.POST("/upload", s.AdminUploadCtrl.Upload, s.AdminService)
	router.POST("/upload/markdown", s.AdminUploadCtrl.MarkdownUpload, s.AdminService)

	// article
	router.POST("/article", s.AdminArticleCtrl.Publish, s.AdminService)
	router.GET("/article/page_:page", s.AdminArticleCtrl.Page, s.AdminService)
	router.GET("/article/:id", s.AdminArticleCtrl.Get, s.AdminService)
	router.GET("/article/count", s.AdminArticleCtrl.Count, s.AdminService)

	// role
	router.GET("/role/permission", s.RoleAuthCtrl.AllPermission, s.AdminService)
	router.Any("/role/init", s.RoleAuthCtrl.InitRole, s.AdminService)
	router.POST("/role", s.RoleAuthCtrl.Save, s.AdminService)
	router.GET("/role/page_:page", s.RoleAuthCtrl.Page, s.AdminService)
	router.GET("/role/:id", s.RoleAuthCtrl.Get, s.AdminService)
	router.GET("/role/all", s.RoleAuthCtrl.All, s.AdminService)

	//router.DELETE("/channel/:id", s.AdminChannelCtrl.Delete)
	//router.GET("/channel/:id", s.AdminChannelCtrl.Get)
	// articcle
	//router.POST("/article/publish", c.AdminArticleCtrl.Publish)
	//router.GET("/article/page_:page", c.AdminArticleCtrl.Page)
	//router.GET("/article/:id", c.AdminArticleCtrl.Get)
	//router.DELETE("/article/:id", c.AdminArticleCtrl.Delete)
	// upload
	//router.POST("/upload", c.AdminUploadCtrl.Upload)
	//router.POST("/upload/base64", c.AdminUploadCtrl.UploadBase64)
	// dict
	//router.GET("/dict/init", c.AdminDictCtrl.Init)
	//router.GET("/dict/all", c.AdminDictCtrl.All)
	//router.PUT("/dict", c.AdminDictCtrl.Edit)
}

func (s *Controller) registerWeb(echo *echo.Echo, prefix string) {
	echo.GET("/", buildHandlerFunc(s.IndexCtrl.Index, s.AdminService))
	echo.GET("/channel/:parent/:title", buildHandlerFunc(s.IndexCtrl.Channel, s.AdminService))
}

func (c *Controller) registerAPI(group *echo.Group, prefix string) {

}

func (s *Controller) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var tmpData map[string]interface{}
	if data == nil {
		tmpData = make(map[string]interface{})
	} else {
		tmpData = data.(map[string]interface{})
	}
	return s.SkinCtrl.Render(w, name, tmpData)
}

type ControllerFunc func(ctx *ctrl.HTTPContext) error

type RouterRegister struct {
	group        *echo.Group
	prefix       string
	buildOption  bool
	renderSource string
}

func (s *RouterRegister) OPTIONS(path string) {
	s.group.OPTIONS(path, func(e echo.Context) error {
		return e.JSON(http.StatusOK, "ok")
	})
}

func (s *RouterRegister) GET(path string, fn ControllerFunc, adminService *service.AdminService) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.GET(path, buildHandlerFunc(fn, adminService))
	fmt.Printf("Register Router[GET], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) POST(path string, fn ControllerFunc, adminService *service.AdminService) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.POST(path, buildHandlerFunc(fn, adminService))
	fmt.Printf("Register Router[POST], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) PUT(path string, fn ControllerFunc, adminService *service.AdminService) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.PUT(path, buildHandlerFunc(fn, adminService))
	fmt.Printf("Register Router[PUT], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) DELETE(path string, fn ControllerFunc, adminService *service.AdminService) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.DELETE(path, buildHandlerFunc(fn, adminService))
	fmt.Printf("Register Router[DELETE], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) Any(path string, fn ControllerFunc, adminService *service.AdminService) {
	s.group.Any(path, buildHandlerFunc(fn, adminService))
	fmt.Printf("Register Router[Any], Path: %v%v\n", s.prefix, path)
}

func buildHandlerFunc(fn ControllerFunc, adminService *service.AdminService) echo.HandlerFunc {
	return func(e echo.Context) error {
		ctx := new(ctrl.HTTPContext)
		ctx.Context = e
		ctx.Result = &ctrl.HTTPResult{Data: make(map[string]interface{})}
		token := e.Request().Header.Get("ACCESS-TOKEN")
		if token != "" {
			admin, _ := adminService.CheckToken(token)
			if admin != nil {
				ctx.Admin = admin
			}
		}
		return fn(ctx)
	}
}
