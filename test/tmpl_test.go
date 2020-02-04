package test

import (
	"github.com/GoodHot/TinyCMS/common/render"
	"os"
	"testing"
)

func TestTmpl1(t *testing.T) {
	r := &render.HTMLRenderer{
		LayoutDir:    "layout",
		ComponentDir: "component",
		TemplateDir:  "../resource/admin",
		Suffix:       ".html",
		Compress:     false,
	}
	r.Init(nil)
	r.Render(os.Stdout, "index", nil)
}
