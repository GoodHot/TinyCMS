package controller

import (
	"github.com/GoodHot/TinyCMS/controller/admin"
	"github.com/labstack/echo"
	"strconv"
)

type Controller struct {
	AdminChannelCtrl *admin.AdminChannelCtrl `ioc:"auto"`
	Port             int                     `val:"${server.port}"`
}

func (c *Controller) StartUp() {
	e := echo.New()
	e.HideBanner = true

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(c.Port)))
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
