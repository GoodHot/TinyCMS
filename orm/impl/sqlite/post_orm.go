package sqlite

import (
	"database/sql"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type PostORMImpl struct {
	DB *datasource.DBPostORM
}

func (orm *PostORMImpl) Save(post *trait.Post) *core.Err {
	err := orm.DB.Insert(&datasource.DBPostModel{
		Title:           post.Title,
		Content:         post.Content,
		ContentHTML:     post.ContentHTML,
		Excerpt:         post.Excerpt,
		Image:           post.Image,
		URL:             post.URL,
		ChannelID:       post.ChannelID,
		TagsID:          post.TagsID,
		Visible:         int(post.Visible),
		Author:          post.Author,
		MetaTitle:       post.MetaTitle,
		MetaDescription: post.MetaDescription,
		Created:         post.Created,
		Modified:        post.Modified,
		PublishTime:     post.PublishTime,
		Status:          int(post.Status),
	}).Exec(func(result sql.Result) {
		last, _ := result.LastInsertId()
		post.ID = int(last)
	})
	if err != nil {
		return core.NewErr(core.Err_Post_Save_Fail)
	}
	return nil
}

func (orm *PostORMImpl) Page(query *trait.Query) (*trait.Page, *core.Err) {
	exp := orm.DB.QueryNoLarge()
	if query.Param != nil {
		cid := query.Param["channel_id"].(int)
		if cid != 0 {
			exp.ChannelIDEq(cid)
		}
		status := query.Param["status"].(int)
		if status != 0 {
			exp.StatusEq(status)
		}
	}
	page := &trait.Page{}
	page.Page = query.Page
	page.Size = query.Size
	count, err := exp.ExecQueryCount()
	if err != nil {
		return nil, core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	exp.Limit(query.Size, (query.Page - 1)*query.Size)
	if query.Order != nil {
		if query.Order["id"] == "desc" {
			exp.OrderByIDDesc()
		}
	}
	rst, err := exp.ExecQuery()
	if err != nil {
		return nil, core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	total := count / page.Size
	if count%page.Size > 0 {
		total++
	}
	page.Count = count
	page.Total = total
	page.List = orm.convert(rst...)
	return page, nil
}

func (*PostORMImpl) convert(model ...*datasource.DBPostModel) []trait.Post {
	var rst []trait.Post
	for _, item := range model {
		rst = append(rst, trait.Post{
			BaseORM: trait.BaseORM{
				ID: item.ID,
				Created: item.Created,
				Modified: item.Modified,
			},
			Title:           item.Title,
			Content:         item.Content,
			ContentHTML:     item.ContentHTML,
			Excerpt:         item.Excerpt,
			Image:           item.Image,
			URL:             item.URL,
			ChannelID:       item.ChannelID,
			TagsID:          item.TagsID,
			Visible:         trait.VisibleLevel(item.Visible),
			PublishTime:     item.PublishTime,
			MetaTitle:       item.MetaTitle,
			MetaDescription: item.MetaDescription,
			Author:          item.Author,
			Status:          trait.PostStatus(item.Status),
		})
	}
	return rst
}
