package controller

import (
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/controller/admin"
)

type Controller struct {
	AdminChannelCtrl *admin.AdminChannelCtrl `ioc:"auto"`
	Config           *config.Config          `ioc:"auto"`
	Name             string                  `val:"${cache.type}"`
	In               int                     `val:"${cache.db_index}"`
	Bo               bool                    `val:"${cache.enable}"`
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
