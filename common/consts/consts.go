package consts

const (
	skinTemplateJsonName = "skin.json"

	// 皮肤模板配置
	configSkinTemplate = "skin.template"
)

// 皮肤模板的配置json文件名
func SkinTemplateJsonName() string {
	return skinTemplateJsonName
}

// 皮肤模板最少包含的文件
func SkinTemplateContainPage() []string {
	return []string{"index.html", "article.html"}
}

func ConfigSkinTemplate() string {
	return configSkinTemplate
}