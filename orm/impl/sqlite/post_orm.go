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

const allPostColumn = "id, title, content, content_html as contenthtml, excerpt, image, url, channel_id as channelid, tags_id as tagsid, visible, author, meta_title metatitle, meta_description metadescription, created, modified, publish_time as publishtime, status"
const allPostColumnNoLarge = "id, title, excerpt, image, url, channel_id as channelid, tags_id as tagsid, visible, author, meta_title metatitle, meta_description metadescription, created, modified, publish_time as publishtime, status"

func (orm *PostORMImpl) Save(post *trait.Post) *core.Err {
	sql := "insert into t_post(title, content, content_html, excerpt, image, url, channel_id,  tags_id, visible, author, meta_title, meta_description, created, modified) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	rst, err := orm.DB.Exec(sql, post.Title, post.Content, post.ContentHTML, post.Excerpt, post.Image, post.URL, post.ChannelID, post.TagsID, post.Visible, post.Author, post.MetaTitle, post.MetaDescription, post.Created, post.Modified, post.PublishTime, post.Status)
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
	sql := "select %v from t_post where 1 = 1 %v %v %v %v %v"
	whereSQL, whereParam := query.BuildWhere()
	lastIDSQL, lastIDParam := query.BuildLastID()
	limitSQL, limitParam := query.BuildLimit()
	orderSQL := query.BuildOrder()
	param = append(param, whereParam...)
	param = append(param, lastIDParam...)
	param = append(param, limitParam...)
	err := orm.DB.Select(&list, fmt.Sprintf(sql, allPostColumnNoLarge, lastIDSQL, whereSQL, "", orderSQL, limitSQL), param...)
	if err != nil {
		return nil, core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	var count int
	err = orm.DB.Get(&count, fmt.Sprintf(sql, "count(1)", "", whereSQL, "", "", ""), param...)
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
