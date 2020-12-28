package sqlite

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type ChannelORMImpl struct {
	DB *sqlx.DB
}

const allColumn = "id, avatar, avatar_svg as avatarsvg, name, path, visible, meta_title as metatitle, meta_description as metadescription, sort, parent_id as parentid"

func (orm ChannelORMImpl) UpdateSort(data []trait.Channel) *core.Err {
	sql := "update t_channel set sort = case id "
	var ids []string
	for index, channel := range data {
		ids = append(ids, strconv.Itoa(channel.ID))
		sql += fmt.Sprintf("when %v then %v ", channel.ID, index)
	}
	sql += " end where id in (" + strings.Join(ids, ",") + ")"
	_, err := orm.DB.Exec(sql)
	if err != nil {
		return core.NewErr(core.Err_Channel_Change_Sort_Fail)
	}
	return nil
}

func (orm ChannelORMImpl) GetAll() (*[]trait.Channel, *core.Err) {
	sql := `select %v from t_channel order by sort asc, id asc`
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
	sql := `INSERT INTO "t_channel"("name", "path", "avatar", "visible", "meta_title", "meta_description", "sort", "parent_id") VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	_, err := orm.DB.Exec(sql, channel.Name, channel.Path, channel.Avatar, channel.Visible, channel.MetaTitle, channel.MetaDescription, channel.Sort, channel.ParentID)
	if err != nil {
		return core.NewErr(core.Err_Channel_Save_Fail)
	}
	return nil
}
