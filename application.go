package main

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server"
)

func main() {
	ioc := new(core.IOC)
	err := ioc.Init(&server.HTTPServer{}, "./config/config.json")
	if err != nil {
		panic(err)
	}
	httpServer := ioc.GetMain().(*server.HTTPServer)
	httpServer.Startup()
}
