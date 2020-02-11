package web

import (
	"github.com/GoodHot/TinyCMS/common/render"
	"html/template"
	"io"
)

type SkinCtrl struct {
	TemplateDir  string `val:"${server.skin_template_dir}"`
	HTMLCompress bool   `val:"${server.html_compress}"`
	templateSkin string
	webRender    *render.HTMLRenderer
}

func (s *SkinCtrl) Render(writer io.Writer, name string, data map[string]interface{}) error {
	if s.webRender == nil {
		s.templateSkin = "default"
		// init template render
		s.webRender = &render.HTMLRenderer{
			LayoutDir:    "layout",
			ComponentDir: "component",
			TemplateDir:  s.TemplateDir + "/" + s.templateSkin,
			Suffix:       ".html",
			Compress:     s.HTMLCompress,
		}
		funcMap := template.FuncMap{
			"noescape": func(s string) template.HTML {
				return template.HTML(s)
			},
		}
		s.webRender.Init(funcMap)
	}
	return s.webRender.Render(writer, name, data)
}
