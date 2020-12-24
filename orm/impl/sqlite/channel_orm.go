package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type ChannelORMImpl struct {
	DB *sqlx.DB
}

func (orm ChannelORMImpl) GetByPath(path string) (*trait.Channel, *core.Err) {
	sql := `
		select id, avatar, name, path, visibility, meta_title as metatitle, meta_description as metadescription, sort, parent_id as parentid
		from t_channel
		where path = ?
		limit 1
	`
	row := orm.DB.QueryRowx(sql, path)
	if row.Err() != nil {
		return nil, core.NewErr(core.Err_Channel_Not_Found_By_Path)
	}
	var channel trait.Channel
	err := row.StructScan(&channel)
	if err != nil {
		return nil, core.NewErr(core.Err_Sys_Server)
	}
	return &channel, nil
}

func (orm ChannelORMImpl) Save(channel *trait.Channel) *core.Err {
	sql := `INSERT INTO "t_channel"("name", "path", "avatar", "visibility", "meta_title", "meta_description", "sort", "parent_id") VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	_, err := orm.DB.Exec(sql, channel.Name, channel.Path, channel.Avatar, channel.Visibility, channel.MetaTitle, channel.MetaDescription, channel.Sort, channel.ParentID)
	if err != nil {
		return core.NewErr(core.Err_Channel_Save_fail)
	}
	return nil
}
