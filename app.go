package main

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/controller"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/spring"
	"os"
)

func main() {
	env := "prod"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	fmt.Println(strs.Fmt("application running with '%s'", env))
	err := spring.Load(strs.Fmt("resource/config/config_%s.json", env))
	if err != nil {
		panic(err)
	}
	spring := spring.AppCtx()

	orm := &orm.ORM{}
	spring.Reg(orm)
	orm.Init(&model.Admin{}, &model.Channel{})

	ctrl := &controller.Controller{}
	spring.Reg(ctrl)
	ctrl.StartUp()
}
