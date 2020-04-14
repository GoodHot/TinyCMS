package render

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TinyRender struct {
	TemplateDir  string
	HTMLCompress bool
	FuncMap      template.FuncMap
	htmlRender   *HTMLRenderer
	templateSkin string
}

func (s *TinyRender) Render(writer io.Writer, name string, data interface{}, ctx echo.Context) error {
	if s.htmlRender == nil {
		s.initHTMLRender()
	}
	return s.htmlRender.Render(writer, name, data)
}

func (s *TinyRender) initHTMLRender() {
	s.templateSkin = "default"
	s.htmlRender = &HTMLRenderer{
		LayoutDir:    "layout",
		ComponentDir: "component",
		TemplateDir:  s.TemplateDir + "/" + s.templateSkin,
		Suffix:       ".html",
		Compress:     s.HTMLCompress,
	}
	s.htmlRender.Init(s.FuncMap)
}
