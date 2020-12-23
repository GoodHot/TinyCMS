package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type ChannelORMImpl struct {
	DB *sqlx.DB
}

func (my ChannelORMImpl) Save(channel *trait.Channel) *core.Err {
	sql := `INSERT INTO "t_channel"("name", "path", "avatar", "visibility", "meta_title", "meta_description", "sort", "parent_id") VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	_, err := my.DB.Exec(sql, channel.Name, channel.Path, channel.Avatar, channel.Visibility, channel.MetaTitle, channel.MetaDescription, channel.Sort, channel.ParentID)
	if err != nil {
		return core.NewErr(core.Err_Channel_Save_fail)
	}
	return nil
}
