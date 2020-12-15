package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Index struct {
	Cache        *common.Cache         `ioc:"auto"`
	AdminService *service.AdminService `ioc:"auto"`
}

func (my *Index) Index(ctx *http.Context) error {
	go time.Sleep(3 * time.Second)
	ctx.Ctx.JSON(200, gin.H{"hello": "world"})
	return nil
}

func (my *Index) List(ctx *http.Context) error {
	my.Cache.Ins().Get("aaa", "bbb")
	fmt.Println("in index ctx")
	return nil
}

type signinForm struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (my *Index) Signin(ctx *http.Context) error {
	var signin signinForm
	if ctx.Bind(&signin) == nil {
		log.Println(signin.Account)
		log.Println(signin.Password)
	}
	return ctx.HTML("hello")
}
