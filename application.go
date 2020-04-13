package main

import (
	"errors"
	"fmt"
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/strs"
	"os"
)

func main() {
	// 判断运行环境
	env := "dev"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	// 根据运行环境读取不同配置文件
	configFile := strs.Fmt("config/config_%s.json", env)
	// 加载配置文件
	brain, err := brain.Load(configFile)
	if err != nil {
		panic(err)
	}
	brain.Reg(&XXX{})
	brain.Reg(&YYY{})
}

type YYY struct{}

type XXX struct{}

func (*XXX) Startup() error {
	fmt.Println("gogogo")
	return errors.New("fffsdf")
}
