package main

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/service/install"
	"github.com/GoodHot/TinyCMS/spring"
	"os"
)

func run(cfg *config.Config) {
	// checking installed
	install := &install.InstallService{}
	if !install.CheckInstalled(cfg.GetDirPath("/app_config/lock/install.lock")) {
		fmt.Println("application not install, waiting init install process")
		// TODO init install process
		return
	}

}

func main() {
	// init configuration
	spring := spring.AppCtx()
	cfg := spring.Reg(&config.Config{}).(*config.Config)

	// get env
	cfg.Env = "prod"
	if len(os.Args) > 1 {
		cfg.Env = os.Args[1]
	}
	fmt.Println(strs.Fmt("application running with '%s'", cfg.Env))

	// get application directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Unable to get execute directory")
		os.Exit(0)
		return
	}
	cfg.AppDir = dir

	// run
	run(cfg)
}
