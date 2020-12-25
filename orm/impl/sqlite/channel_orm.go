package sqlite

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type ChannelORMImpl struct {
	DB *sqlx.DB
}

const allColumn = "id, avatar, avatar_svg as avatarsvg, name, path, visibility, meta_title as metatitle, meta_description as metadescription, sort, parent_id as parentid"

func (orm ChannelORMImpl) GetAll() (*[]trait.Channel, *core.Err) {
	sql := `select %v from t_channel`
	var cs []trait.Channel
	err := orm.DB.Select(&cs, fmt.Sprintf(sql, allColumn))
	if err != nil {
		return nil, core.NewErr(core.Err_Channel_Can_Not_Get_List)
	}
	return &cs, nil
}

func (orm ChannelORMImpl) GetByPath(path string) (*trait.Channel, *core.Err) {
	sql := `
		select %v
		from t_channel
		where path = ?
		limit 1
	`
	var channel trait.Channel
	err := orm.DB.Get(&channel, fmt.Sprintf(sql, allColumn), path)
	if err != nil {
		return nil, core.NewErr(core.Err_Channel_Not_Found_By_Path)
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
