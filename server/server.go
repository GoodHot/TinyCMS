package server

import (
	"github.com/GoodHot/TinyCMS/server/router"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	ServerAddr string         `val:"${server.addr}"`
	Router     *router.Router `ioc:"auto"`
}

// 启动服务
func (server *HTTPServer) Startup() {
	engine := gin.Default()
	server.Router.Register(engine.Group("/"))
	engine.Run(server.ServerAddr)
}
