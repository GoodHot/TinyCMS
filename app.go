package main

import (
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/spring"
)

func run(cfg *config.Config) {
	// checking installed
	//install := &install.InstallService{}
	//if !install.CheckInstalled(cfg.GetDirPath(cfg.ConfigDir + "/lock/install.lock")) {
	//	fmt.Println("application not install, waiting init install process")
	//	// TODO init install process
	//	return
	//}
	//cfg.Load()
	//spring := spring.AppCtx()
	//ctrl := spring.Reg(&controller.Controller{}).(*controller.Controller)
	//ctrl.StartUp()
}

func main() {
	err := spring.Load("app_config/app_dev.json")
	if err != nil {
		panic(err)
	}
	spring := spring.AppCtx()
	// init configuration
	//spring := spring.AppCtx()

	//cfg := spring.Reg(&config.Config{}).(*config.Config)
	//
	//// get env
	//cfg.Env = "prod"
	//if len(os.Args) > 1 {
	//	cfg.Env = os.Args[1]
	//}
	//fmt.Println(strs.Fmt("application running with '%s'", cfg.Env))
	//
	//// get application directory
	//dir, err := os.Getwd()
	//if err != nil {
	//	fmt.Errorf("Unable to get execute directory")
	//	os.Exit(0)
	//	return
	//}
	//cfg.AppDir = dir
	//cfg.ConfigDir = "/app_config"
	//// run
	//run(cfg)
}
