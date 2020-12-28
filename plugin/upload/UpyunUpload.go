package upload

import (
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/service"
)

type UpyunUploadPlugin struct {
}

func (*UpyunUploadPlugin) Install(param *service.ParamOperator) {
	param.Set(&trait.PluginParam{
		Key:         "Bucket",
		Value:       "",
		Description: "Upyun Bucket",
		CanEdit:     true,
		Visible:     true,
		ValueType:   trait.FormType_Input,
	})
	param.Set(&trait.PluginParam{
		Key:         "Operator",
		Value:       "",
		Description: "Upyun Operator",
		CanEdit:     true,
		Visible:     true,
		ValueType:   trait.FormType_Input,
	})
	param.Set(&trait.PluginParam{
		Key:         "Password",
		Value:       "",
		Description: "Upyun Password",
		CanEdit:     true,
		Visible:     true,
		ValueType:   trait.FormType_Input,
	})
}

