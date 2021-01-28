package server

import (
	"github.com/GoodHot/TinyCMS/server/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type HTTPServer struct {
	Env           string                `val:"${env}"`
	ServerAddr    string                `val:"${server.addr}"`
	Static        string                `val:"${server.static}"`
	Router        *router.Router        `ioc:"auto"`
}

// 启动服务
func (server *HTTPServer) Startup() {
	engine := echo.New()
	engine.Static(server.Static, server.Static)
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
	server.Router.Register(engine, engine.Group(""))
	engine.Logger.Fatal(engine.Start(server.ServerAddr))
}
