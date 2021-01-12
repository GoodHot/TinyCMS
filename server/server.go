package server

import (
	"github.com/GoodHot/TinyCMS/plugin"
	"github.com/GoodHot/TinyCMS/server/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	Env           string                `val:"${env}"`
	ServerAddr    string                `val:"${server.addr}"`
	Router        *router.Router        `ioc:"auto"`
	PluginManager *plugin.PluginManager `ioc:"auto"`
}

// 启动服务
func (server *HTTPServer) Startup() {
	engine := echo.New()
	if server.Env == "dev" || server.Env == "test" {
		engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		}))
		engine.OPTIONS("*", func(context echo.Context) error {
			return nil
		})
	}
	server.PluginManager.Mount()
	server.Router.Register(engine.Group(""))
	engine.Logger.Fatal(engine.Start(server.ServerAddr))
}
