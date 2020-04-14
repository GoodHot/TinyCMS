package controller

import (
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/render"
	"github.com/GoodHot/TinyCMS/controller/admin"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Controller struct {
	Log                brain.Logger          `ioc:"auto"`
	AdminAuthCtrl      *admin.AdminAuthCtrl  `ioc:"auto"`
	AdminService       *service.AdminService `ioc:"auto"`
	ServerStatic       string                `val:"${server.static}"`            // 静态文件夹
	ServerPort         int                   `val:"${server.port}"`              // 服务启动端口
	ServerTemplateDir  string                `val:"${server.skin_template_dir}"` // 模板文件夹
	ServerHTMLCompress bool                  `val:"${server.html_compress}"`     // HTML代码是否压缩
	AdminPrefix        string                `val:"${server.admin_prefix}"`      // Admin接口访问路径
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
	s.registerAdmin(e.Group(s.AdminPrefix), s.AdminPrefix)
	// 启动服务
	return e.Start(":" + strconv.Itoa(s.ServerPort))
}

func (s *Controller) registerAdmin(group *echo.Group, prefix string) {
	router := &RouterRegister{
		Group:        group,
		Prefix:       prefix,
		BuildOption:  false,
		AdminService: s.AdminService,
		Log:          s.Log,
	}
	router.POST("/auth/login", s.AdminAuthCtrl.Login)
}
