package sqlite

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type PostORMImpl struct {
	DB *sqlx.DB
}

const allPostColumn = "id, title, content, content_html as contenthtml, tags_id as tagsid, channel_id as channelid, visible, author, created, modified"
const allPostColumnNoLarge = "id, title, tags_id as tagsid, channel_id as channelid, visible, author, created, modified"

func (orm *PostORMImpl) Save(post *trait.Post) *core.Err {
	sql := "insert into t_post(title, content, content_html, tags_id, channel_id, visible, author, created, modified) values(?,?,?,?,?,?,?,?,?)"
	rst, err := orm.DB.Exec(sql, post.Title, post.Content, post.ContentHTML, post.TagsID, post.ChannelID, post.Visible, post.Author, post.Created, post.Modified)
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

func (orm *PostORMImpl) Page(baseQuery *trait.Query) (*trait.Page, *core.Err) {
	query := ToSQLiteQuery(baseQuery)
	var page trait.Page
	var list []trait.Post
	var param []interface{}
	sql := "select %v from t_post where 1 = 1 %v %v %v %v"
	lastIDSQL, lastIDParam := query.BuildLastID()
	limitSQL, limitParam := query.BuildLimit()
	orderSQL := query.BuildOrder()
	param = append(param, lastIDParam...)
	param = append(param, limitParam...)
	err := orm.DB.Select(&list, fmt.Sprintf(sql, allPostColumnNoLarge, lastIDSQL, "", orderSQL, limitSQL), param...)
	if err != nil {
		return nil, core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	var count int
	err = orm.DB.Get(&count, fmt.Sprintf(sql, "count(1)", "", "", "", ""))
	if err != nil {
		return nil, core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	total := count / query.Size
	if count%query.Size > 0 {
		total++
	}
	page.Page = query.Page
	page.Size = query.Size
	page.Count = count
	page.Total = total
	page.List = &list
	return &page, nil
}
