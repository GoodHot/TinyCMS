package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"strconv"
	"strings"
)

type ChannelORMImpl struct {
	DB *datasource.DBChannelORM
}

func (orm *ChannelORMImpl) UpdateSort(data []trait.Channel) *core.Err {
	sql := "update t_channel set sort = case id "
	var ids []string
	for index, channel := range data {
		ids = append(ids, strconv.Itoa(channel.ID))
		sql += fmt.Sprintf("when %v then %v ", channel.ID, index)
	}
	sql += " end where id in (" + strings.Join(ids, ",") + ")"
	_, err := orm.DB.DB.Exec(sql)
	if err != nil {
		return core.NewErr(core.Err_Channel_Change_Sort_Fail)
	}
	return nil
}

func (orm *ChannelORMImpl) GetAll() ([]trait.Channel, *core.Err) {
	results, err := orm.DB.Query().OrderBySortAsc().OrderByIDAsc().ExecQuery()
	if err != nil {
		return nil, core.NewErr(core.Err_Channel_Can_Not_Get_List)
	}
	return orm.convert(results...), nil
}

func (orm *ChannelORMImpl) GetByPath(path string) (*trait.Channel, *core.Err) {
	model, err := orm.DB.Query().PathEq(path).Limit(1, 0).ExecQueryOne()
	if err != nil {
		return nil, core.NewErr(core.Err_Channel_Not_Found_By_Path)
	}
	return &orm.convert(model)[0], nil
}

func (orm *ChannelORMImpl) Save(channel *trait.Channel) *core.Err {
	err := orm.DB.Insert(&datasource.DBChannelModel{
		Name:            channel.Name,
		Path:            channel.Path,
		AvatarSVG:       channel.AvatarSVG,
		Avatar:          channel.Avatar,
		Visible:         int(channel.Visible),
		MetaTitle:       channel.MetaTitle,
		MetaDescription: channel.MetaDescription,
		Sort:            channel.Sort,
		ParentID:        channel.ParentID,
	}).Exec(func(result sql.Result) {
		last, _ := result.LastInsertId()
		channel.ID = int(last)
	})
	if err != nil {
		return core.NewErr(core.Err_Channel_Save_Fail)
	}
	return nil
}

func (orm *ChannelORMImpl) convert(model ...*datasource.DBChannelModel) []trait.Channel {
	var rst []trait.Channel
	for _, item := range model {
		rst = append(rst, trait.Channel{
			BaseORM: trait.BaseORM{
				ID: item.ID,
			},
			Avatar:          item.Avatar,
			AvatarSVG:       item.AvatarSVG,
			Name:            item.Name,
			Path:            item.Path,
			Visible:         trait.VisibleLevel(item.Visible),
			MetaTitle:       item.MetaTitle,
			MetaDescription: item.MetaDescription,
			Sort:            item.Sort,
			ParentID:        item.ParentID,
		})
	}
	return rst
}
