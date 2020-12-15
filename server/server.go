package server

import (
	"github.com/GoodHot/TinyCMS/server/router"
	"github.com/labstack/echo"
)

type HTTPServer struct {
	ServerAddr string         `val:"${server.addr}"`
	Router     *router.Router `ioc:"auto"`
}

// 启动服务
func (server *HTTPServer) Startup() {
	engine := echo.New()
	server.Router.Register(engine.Group(""))
	engine.Logger.Fatal(engine.Start(server.ServerAddr))
}
