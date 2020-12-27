package upload

import (
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/service"
)

type LocalUploadPlugin struct {
}

func (*LocalUploadPlugin) Install(param *service.ParamOperator) {
	param.Set(&trait.PluginParam{
		Key:         "prefix",
		Value:       "/files",
		Description: "文件上传前缀",
		CanEdit:     true,
		Visible:     true,
		ValueType:   trait.PluginParamType_Input,
	})
}
