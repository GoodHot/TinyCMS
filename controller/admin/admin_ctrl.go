package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/ctrl"
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
	return ctx.ResultOK()
}
