package main

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/controller"
	"github.com/GoodHot/TinyCMS/spring"
	"os"
)

func main() {
	env := "prod"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	fmt.Println(strs.Fmt("application running with '%s'", env))
	err := spring.Load(strs.Fmt("app_config/config_%s.json", env))
	if err != nil {
		panic(err)
	}
	spring := spring.AppCtx()
	ctrl := &controller.Controller{}
	spring.Reg(ctrl)
	ctrl.StartUp()
}
