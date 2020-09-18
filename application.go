package main

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server"
)

func main() {
	ioc := new(core.IOC)
	ioc.Init(&server.HTTPServer{}, "./config/config.json")
	httpServer := ioc.GetMain().(*server.HTTPServer)
	httpServer.Startup()
}
