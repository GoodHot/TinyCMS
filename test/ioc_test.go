package test

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/controller"
	"github.com/GoodHot/TinyCMS/spring"
	"testing"
)

func TestIOC(t *testing.T) {
	ctx := spring.AppCtx()
	ctrl := &controller.Controller{}
	ctx.Reg(ctrl)
	fmt.Println(ctrl.AdminChannelCtrl.Age)
}
