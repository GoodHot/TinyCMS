package controller

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/common/render"
	"github.com/GoodHot/TinyCMS/controller/admin"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"strconv"
)

type Controller struct {
	AdminChannelCtrl *admin.AdminChannelCtrl `ioc:"auto"`
	Port             int                     `val:"${server.port}"`
	AdminPrefix      string                  `val:"${server.admin_prefix}"`
	WebPrefix        string                  `val:"${server.web_prefix}"`
	APIPrefix        string                  `val:"${server.api_prefix}"`
	Static           string                  `val:"${server.static}"`
	HTMLCompress     bool                    `val:"${server.html_compress}"`
	adminRender      *render.HTMLRenderer
	webRender        *render.HTMLRenderer
}

func (c *Controller) StartUp() {
	e := echo.New()
	e.HideBanner = true

	e.Static(c.Static, c.Static)
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	c.registerAdmin(e.Group(c.AdminPrefix), c.AdminPrefix)
	c.registerWeb(e.Group(c.WebPrefix), c.WebPrefix)
	c.registerAPI(e.Group(c.APIPrefix), c.APIPrefix)

	err := e.Start(":" + strconv.Itoa(c.Port))
	if err != nil {
		panic(err)
	}
}

//func (s *Controller) initRender() {
//	s.adminRender = &render.HTMLRenderer{
//		LayoutDir:    "layout",
//		ComponentDir: "component",
//		TemplateDir:  "resource/admin",
//		Suffix:       ".html",
//		Compress:     s.HTMLCompress,
//	}
//	s.webRender = &render.HTMLRenderer{
//		LayoutDir:    "layout",
//		ComponentDir: "component",
//		TemplateDir:  "resource/skin/default",
//		Suffix:       ".html",
//		Compress:     s.HTMLCompress,
//	}
//}

func (s *Controller) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	return s.adminRender.Render(w, name, data)
}

func (c *Controller) registerAdmin(group *echo.Group, prefix string) {
	router := &RouterRegister{
		group:       group,
		prefix:      prefix,
		buildOption: true,
	}
	router.GET("/hello", c.AdminChannelCtrl.List)
}

func (c *Controller) registerWeb(group *echo.Group, prefix string) {

}

func (c *Controller) registerAPI(group *echo.Group, prefix string) {

}

type ControllerFunc func(ctx *ctrl.HTTPContext) error

type RouterRegister struct {
	group       *echo.Group
	prefix      string
	buildOption bool
}

func (s *RouterRegister) OPTIONS(path string) {
	s.group.OPTIONS(path, func(e echo.Context) error {
		return e.JSON(http.StatusOK, "ok")
	})
}

func (s *RouterRegister) GET(path string, fn ControllerFunc) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.GET(path, buildHandlerFunc(fn))
	fmt.Printf("Register Router[GET], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) POST(path string, fn ControllerFunc) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.POST(path, buildHandlerFunc(fn))
	fmt.Printf("Register Router[POST], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) PUT(path string, fn ControllerFunc) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.PUT(path, buildHandlerFunc(fn))
	fmt.Printf("Register Router[PUT], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) DELETE(path string, fn ControllerFunc) {
	if s.buildOption {
		s.OPTIONS(path)
	}
	s.group.DELETE(path, buildHandlerFunc(fn))
	fmt.Printf("Register Router[DELETE], Path: %v%v\n", s.prefix, path)
}

func (s *RouterRegister) Any(path string, fn ControllerFunc) {
	s.group.Any(path, buildHandlerFunc(fn))
	fmt.Printf("Register Router[Any], Path: %v%v\n", s.prefix, path)
}

func buildHandlerFunc(fn ControllerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		ctx := new(ctrl.HTTPContext)
		ctx.Context = e
		ctx.Result = &ctrl.HTTPResult{Data: make(map[string]interface{})}
		return fn(ctx)
	}
}
