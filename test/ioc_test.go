package test

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/controller"
	"github.com/GoodHot/TinyCMS/spring"
	"regexp"
	"testing"
)

func TestIOC(t *testing.T) {
	ctx := spring.AppCtx()
	ctrl := &controller.Controller{}
	ctx.Reg(ctrl)
	//fmt.Println(ctrl.AdminChannelCtrl.Age)
}


func TestIOC2(t *testing.T){
	err := spring.Load("../app_config/config_dev.json")
	if err != nil {
		panic(err)
	}
	spring := spring.AppCtx()
	ctrl := &controller.Controller{}
	spring.Reg(ctrl)
	//fmt.Println(ctrl.Name)
	//fmt.Println(ctrl.In)
	//fmt.Println(ctrl.Bo)
}

func TestInject(t *testing.T){
	key := "${dd.dfq}"
	compile := regexp.MustCompile("^\\$\\{(.+?)\\}$")
	fmt.Println(compile.MatchString(key))
	fmt.Println(compile.FindAllStringSubmatch(key, -1))
}