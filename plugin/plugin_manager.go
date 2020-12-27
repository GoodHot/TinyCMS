package plugin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/plugin/upload"
	"github.com/GoodHot/TinyCMS/service"
)

type PluginManager struct {
	PluginService *service.PluginService `ioc:"auto"`
}

func (pm *PluginManager) Mount() error {
	err := pm.PluginService.MountInternal(&trait.Plugin{
		Name:        "LocalUpload",
		Version:     "1.0.0",
		Description: "upload file by local",
		Type:        trait.PluginTypeUpload,
	}, &upload.LocalUploadPlugin{})
	if err != nil && err.ErrType == core.Err_Plugin_Exists {
		fmt.Println("这个插件已经存在了")
	}
	return nil
}
