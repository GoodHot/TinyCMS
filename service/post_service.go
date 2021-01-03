package service

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/davidscottmills/goeditorjs"
)

type PostService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (ps *PostService) Save(post *trait.Post) *core.Err {
	htmlEngine := goeditorjs.NewHTMLEngine()
	htmlEngine.RegisterBlockHandlers(
		&goeditorjs.HeaderHandler{},
		&goeditorjs.ParagraphHandler{},
		&goeditorjs.ListHandler{},
		&goeditorjs.CodeBoxHandler{},
	)
	html, err := htmlEngine.GenerateHTML(post.Content)
	if err != nil {
		return core.NewErr(core.Err_Post_Content_Format_Error)
	}
	post.ContentHTML = html
	return ps.ORM.Post.Save(post)
}
