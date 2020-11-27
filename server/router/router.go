package router

import (
	"github.com/GoodHot/TinyCMS/server/router/admin"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/gin-gonic/gin"
)

type Router struct {
	RouterAdmin string       `val:"${props.router.admin}"`
	RouterWeb   string       `val:"${props.router.web}"`
	AdminIndex  *admin.Index `ioc:"auto"`
}

func (router *Router) Register(group *gin.RouterGroup) {
	// 注册admin路由
	adminRouter := group.Group(router.RouterAdmin)
	router.registerAdmin(adminRouter)
	// 注册web路由
	webRouter := group.Group(router.RouterWeb)
	router.registerWeb(webRouter)
}

func (router *Router) registerAdmin(group *gin.RouterGroup) {
	register := http.RouterRegister{Group: group}
	register.GET("/index", router.AdminIndex.Index)
}

func (*Router) registerWeb(router *gin.RouterGroup) {
}
