package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type PostORMImpl struct {
	DB *sqlx.DB
}

func (orm *PostORMImpl) Save(post *trait.Post) *core.Err {
	sql := "insert into t_post(title, content, content_html, tags_id, channel_id, visible) values(?,?,?,?,?,?)"
	rst, err := orm.DB.Exec(sql, post.Title, post.Content, post.ContentHTML, post.TagsID, post.ChannelID, post.Visible)
	if err != nil {
		return core.NewErr(core.Err_Post_Save_Fail)
	}
	id, err := rst.LastInsertId()
	if err != nil {
		return nil
	}
	post.ID = int(id)
	return nil
}
