package admin

import "github.com/GoodHot/TinyCMS/common/ctrl"

type AdminAuthCtrl struct {
}

func (*AdminAuthCtrl) Login(ctx *ctrl.HTTPContext) error {
	return ctx.HTML("")
}
