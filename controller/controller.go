package controller

import (
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/controller/admin"
)

type Controller struct {
	adminChannelCtrl *admin.AdminChannelCtrl `ioc:auto`
	config           *config.Config          `ioc:"auto"`
}

func (c *Controller) StartUp() {
	//e := echo.New()
	//e.Group(c.config.Controller.AdminPrefix, func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	//	return nil
	//})
	//e.Group(c.config.Controller.WebPrefix, func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	//	return nil
	//})
	//e.Group(c.config.Controller.APIPrefix, func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	//	return nil
	//})
}
