package main

import (
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/logger"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/controller"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"os"
)

var BrainContext brain.Brain

func main() {
	// 判断运行环境
	env := "dev"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	// 根据运行环境读取不同配置文件
	configFile := strs.Fmt("config/config_%s.json", env)
	// 加载配置文件
	BrainContext, err := brain.Load(configFile)
	if err != nil {
		panic(err)
	}
	// 设置全局日志
	BrainContext.SetLogger(&logger.TinyLogger{})
	// 创建ORM和数据库表
	tables := []interface{}{
		&model.Admin{},
		&model.Skin{},
		&model.Config{},
		&model.Article{},
		&model.Category{},
		&model.Upload{},
	}
	BrainContext.Register(&orm.ORM{}).(*orm.ORM).AutoMigrate(tables...)
	// 启动controller
	BrainContext.Register(&controller.Controller{})
}
