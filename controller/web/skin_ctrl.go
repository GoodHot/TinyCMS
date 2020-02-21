package web

import (
	"github.com/GoodHot/TinyCMS/common/render"
	"github.com/GoodHot/TinyCMS/common/times"
	"github.com/GoodHot/TinyCMS/service"
	"html/template"
	"io"
	"strings"
	"time"
)

type SkinCtrl struct {
	DictService    *service.DictService    `ioc:"auto"`
	CategoryService *service.CategoryService `ioc:"auto"`
	TemplateDir    string                  `val:"${server.skin_template_dir}"`
	HTMLCompress   bool                    `val:"${server.html_compress}"`
	templateSkin   string
	webRender      *render.HTMLRenderer
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
			"timeFmt": func(t time.Time) string {
				return times.TimeFmt(t, "2006-01-02 15:04")
			},
			"dict": func(key string) string {
				dict := s.DictService.GetByKey(key)
				if dict == nil {
					return ""
				}
				return dict.Value
			},
			//"channelTree": func() []*service.ChannelTree {
			//	channels := s.ChannelService.Tree()
			//	return channels
			//},
			//"getChannel": func(cid uint) *model.Channel {
			//	channel, err := s.ChannelService.Get(int(cid))
			//	if err != nil {
			//		return &model.Channel{}
			//	}
			//	return channel
			//},
			"split": func(str string) []string {
				return strings.Split(str, ",")
			},
		}
		s.webRender.Init(funcMap)
	}
	return s.webRender.Render(writer, name, data)
}
