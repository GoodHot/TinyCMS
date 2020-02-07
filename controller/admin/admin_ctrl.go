package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
)

type AdminAuthCtrl struct {
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*AdminAuthCtrl) Login(ctx *ctrl.HTTPContext) error {
	u := new(user)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	fmt.Println(u)
	admin := model.Admin{}
	admin.Avatar = "http://wx2.sinaimg.cn/mw600/6b465dably1g9jos0m98ej20au0nsgmp.jpg"
	admin.Token = "test"
	ctx.Put("info", admin)
	return ctx.ResultOK()
}

func (*AdminAuthCtrl) Info(ctx *ctrl.HTTPContext) error {
	admin := model.Admin{}
	admin.Avatar = "http://wx2.sinaimg.cn/mw600/6b465dably1g9jos0m98ej20au0nsgmp.jpg"
	admin.Token = "test"
	ctx.Put("info", admin)
	return ctx.ResultOK()
}
