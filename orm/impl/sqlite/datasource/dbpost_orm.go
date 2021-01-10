
package datasource

import (
    "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type DBPostORMExample struct {
	showSQL bool
	db      *sqlx.DB
	sql     string
	where   string
	orderBy string
	limit   string
	param   []interface{}
}

func (exp *DBPostORMExample) And() *DBPostORMExample {
	exp.where += " and"
	return exp
}
func (exp *DBPostORMExample) Or() *DBPostORMExample {
	exp.where += " or"
	return exp
}
func (exp *DBPostORMExample) Exec(handler ...func(result sql.Result)) error {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }

    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	result, err := exp.db.Exec(sql, exp.param...)
	if err != nil {
	    return err
	}
    if len(handler) > 0 {
        for _, h := range handler {
            h(result)
        }
    }
	return nil
}
func (exp *DBPostORMExample) ExecQuery() ([]*DBPostModel, error) {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }

    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	rows, err := exp.db.Query(sql, exp.param...)
	if err != nil {
		return nil, err
	}
	var result []*DBPostModel
	for rows.Next() {
		var model DBPostModel
        if err := rows.Scan( &model.ID, &model.Title, &model.Content, &model.ContentHTML, &model.Excerpt, &model.Image, &model.URL, &model.ChannelID, &model.TagsID, &model.Visible, &model.Author, &model.MetaTitle, &model.MetaDescription, &model.Created, &model.Modified, &model.PublishTime, &model.Status); err != nil {
            return result, err
        }
        result = append(result, &model)
	}
	return result, nil
}
func (exp *DBPostORMExample) ExecQueryOne() (*DBPostModel, error) {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }

    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	rows, err := exp.db.Query(sql, exp.param...)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
	    var model DBPostModel
	    if err := rows.Scan( &model.ID, &model.Title, &model.Content, &model.ContentHTML, &model.Excerpt, &model.Image, &model.URL, &model.ChannelID, &model.TagsID, &model.Visible, &model.Author, &model.MetaTitle, &model.MetaDescription, &model.Created, &model.Modified, &model.PublishTime, &model.Status); err != nil {
            return nil, err
        }
        return &model, nil
	}
	return nil, nil
}
func (exp *DBPostORMExample) ExecPage() error {
	sql := exp.sql
	if exp.where != "" {
		sql += " where " + exp.where
	}
	fmt.Println(sql)
	return nil
}
func (exp *DBPostORMExample) AddParam(param interface{}) *DBPostORMExample {
	exp.param = append(exp.param, param)
	return exp
}

func (exp *DBPostORMExample) Limit(limit int, offset int) *DBPostORMExample {
	exp.limit = " limit ? offset ? "
	exp.param = append(exp.param, limit)
	exp.param = append(exp.param, offset)
	return exp
}




func (exp *DBPostORMExample) IDEq(ID int) *DBPostORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBPostORMExample) IDGt(ID int) *DBPostORMExample {
	exp.where += "id > ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBPostORMExample) IDLt(ID int) *DBPostORMExample {
	exp.where += "id < ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBPostORMExample) IDEqGt(ID int) *DBPostORMExample {
	exp.where += "id >= ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBPostORMExample) IDEqLt(ID int) *DBPostORMExample {
	exp.where += "id <= ?"
	exp.param = append(exp.param, ID)
	return exp
}







func (exp *DBPostORMExample) TitleEq(Title string) *DBPostORMExample {
	exp.where += "title = ?"
	exp.param = append(exp.param, Title)
	return exp
}

func (exp *DBPostORMExample) TitleLike(Title string) *DBPostORMExample {
	exp.where += "title like '%' || ? || '%'"
	exp.param = append(exp.param, Title)
	return exp
}

func (exp *DBPostORMExample) TitleLikeAfter(Title string) *DBPostORMExample {
	exp.where += "title like ? || '%'"
	exp.param = append(exp.param, Title)
	return exp
}

func (exp *DBPostORMExample) TitleLikeBefore(Title string) *DBPostORMExample {
	exp.where += "title like  '%' || ?"
	exp.param = append(exp.param, Title)
	return exp
}





func (exp *DBPostORMExample) ContentEq(Content string) *DBPostORMExample {
	exp.where += "content = ?"
	exp.param = append(exp.param, Content)
	return exp
}

func (exp *DBPostORMExample) ContentLike(Content string) *DBPostORMExample {
	exp.where += "content like '%' || ? || '%'"
	exp.param = append(exp.param, Content)
	return exp
}

func (exp *DBPostORMExample) ContentLikeAfter(Content string) *DBPostORMExample {
	exp.where += "content like ? || '%'"
	exp.param = append(exp.param, Content)
	return exp
}

func (exp *DBPostORMExample) ContentLikeBefore(Content string) *DBPostORMExample {
	exp.where += "content like  '%' || ?"
	exp.param = append(exp.param, Content)
	return exp
}





func (exp *DBPostORMExample) ContentHTMLEq(ContentHTML string) *DBPostORMExample {
	exp.where += "content_html = ?"
	exp.param = append(exp.param, ContentHTML)
	return exp
}

func (exp *DBPostORMExample) ContentHTMLLike(ContentHTML string) *DBPostORMExample {
	exp.where += "content_html like '%' || ? || '%'"
	exp.param = append(exp.param, ContentHTML)
	return exp
}

func (exp *DBPostORMExample) ContentHTMLLikeAfter(ContentHTML string) *DBPostORMExample {
	exp.where += "content_html like ? || '%'"
	exp.param = append(exp.param, ContentHTML)
	return exp
}

func (exp *DBPostORMExample) ContentHTMLLikeBefore(ContentHTML string) *DBPostORMExample {
	exp.where += "content_html like  '%' || ?"
	exp.param = append(exp.param, ContentHTML)
	return exp
}





func (exp *DBPostORMExample) ExcerptEq(Excerpt string) *DBPostORMExample {
	exp.where += "excerpt = ?"
	exp.param = append(exp.param, Excerpt)
	return exp
}

func (exp *DBPostORMExample) ExcerptLike(Excerpt string) *DBPostORMExample {
	exp.where += "excerpt like '%' || ? || '%'"
	exp.param = append(exp.param, Excerpt)
	return exp
}

func (exp *DBPostORMExample) ExcerptLikeAfter(Excerpt string) *DBPostORMExample {
	exp.where += "excerpt like ? || '%'"
	exp.param = append(exp.param, Excerpt)
	return exp
}

func (exp *DBPostORMExample) ExcerptLikeBefore(Excerpt string) *DBPostORMExample {
	exp.where += "excerpt like  '%' || ?"
	exp.param = append(exp.param, Excerpt)
	return exp
}





func (exp *DBPostORMExample) ImageEq(Image string) *DBPostORMExample {
	exp.where += "image = ?"
	exp.param = append(exp.param, Image)
	return exp
}

func (exp *DBPostORMExample) ImageLike(Image string) *DBPostORMExample {
	exp.where += "image like '%' || ? || '%'"
	exp.param = append(exp.param, Image)
	return exp
}

func (exp *DBPostORMExample) ImageLikeAfter(Image string) *DBPostORMExample {
	exp.where += "image like ? || '%'"
	exp.param = append(exp.param, Image)
	return exp
}

func (exp *DBPostORMExample) ImageLikeBefore(Image string) *DBPostORMExample {
	exp.where += "image like  '%' || ?"
	exp.param = append(exp.param, Image)
	return exp
}





func (exp *DBPostORMExample) URLEq(URL string) *DBPostORMExample {
	exp.where += "url = ?"
	exp.param = append(exp.param, URL)
	return exp
}

func (exp *DBPostORMExample) URLLike(URL string) *DBPostORMExample {
	exp.where += "url like '%' || ? || '%'"
	exp.param = append(exp.param, URL)
	return exp
}

func (exp *DBPostORMExample) URLLikeAfter(URL string) *DBPostORMExample {
	exp.where += "url like ? || '%'"
	exp.param = append(exp.param, URL)
	return exp
}

func (exp *DBPostORMExample) URLLikeBefore(URL string) *DBPostORMExample {
	exp.where += "url like  '%' || ?"
	exp.param = append(exp.param, URL)
	return exp
}



func (exp *DBPostORMExample) ChannelIDEq(ChannelID int) *DBPostORMExample {
	exp.where += "channel_id = ?"
	exp.param = append(exp.param, ChannelID)
	return exp
}
func (exp *DBPostORMExample) ChannelIDGt(ChannelID int) *DBPostORMExample {
	exp.where += "channel_id > ?"
	exp.param = append(exp.param, ChannelID)
	return exp
}
func (exp *DBPostORMExample) ChannelIDLt(ChannelID int) *DBPostORMExample {
	exp.where += "channel_id < ?"
	exp.param = append(exp.param, ChannelID)
	return exp
}
func (exp *DBPostORMExample) ChannelIDEqGt(ChannelID int) *DBPostORMExample {
	exp.where += "channel_id >= ?"
	exp.param = append(exp.param, ChannelID)
	return exp
}
func (exp *DBPostORMExample) ChannelIDEqLt(ChannelID int) *DBPostORMExample {
	exp.where += "channel_id <= ?"
	exp.param = append(exp.param, ChannelID)
	return exp
}







func (exp *DBPostORMExample) TagsIDEq(TagsID string) *DBPostORMExample {
	exp.where += "tags_id = ?"
	exp.param = append(exp.param, TagsID)
	return exp
}

func (exp *DBPostORMExample) TagsIDLike(TagsID string) *DBPostORMExample {
	exp.where += "tags_id like '%' || ? || '%'"
	exp.param = append(exp.param, TagsID)
	return exp
}

func (exp *DBPostORMExample) TagsIDLikeAfter(TagsID string) *DBPostORMExample {
	exp.where += "tags_id like ? || '%'"
	exp.param = append(exp.param, TagsID)
	return exp
}

func (exp *DBPostORMExample) TagsIDLikeBefore(TagsID string) *DBPostORMExample {
	exp.where += "tags_id like  '%' || ?"
	exp.param = append(exp.param, TagsID)
	return exp
}



func (exp *DBPostORMExample) VisibleEq(Visible int) *DBPostORMExample {
	exp.where += "visible = ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBPostORMExample) VisibleGt(Visible int) *DBPostORMExample {
	exp.where += "visible > ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBPostORMExample) VisibleLt(Visible int) *DBPostORMExample {
	exp.where += "visible < ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBPostORMExample) VisibleEqGt(Visible int) *DBPostORMExample {
	exp.where += "visible >= ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBPostORMExample) VisibleEqLt(Visible int) *DBPostORMExample {
	exp.where += "visible <= ?"
	exp.param = append(exp.param, Visible)
	return exp
}





func (exp *DBPostORMExample) AuthorEq(Author int) *DBPostORMExample {
	exp.where += "author = ?"
	exp.param = append(exp.param, Author)
	return exp
}
func (exp *DBPostORMExample) AuthorGt(Author int) *DBPostORMExample {
	exp.where += "author > ?"
	exp.param = append(exp.param, Author)
	return exp
}
func (exp *DBPostORMExample) AuthorLt(Author int) *DBPostORMExample {
	exp.where += "author < ?"
	exp.param = append(exp.param, Author)
	return exp
}
func (exp *DBPostORMExample) AuthorEqGt(Author int) *DBPostORMExample {
	exp.where += "author >= ?"
	exp.param = append(exp.param, Author)
	return exp
}
func (exp *DBPostORMExample) AuthorEqLt(Author int) *DBPostORMExample {
	exp.where += "author <= ?"
	exp.param = append(exp.param, Author)
	return exp
}







func (exp *DBPostORMExample) MetaTitleEq(MetaTitle string) *DBPostORMExample {
	exp.where += "meta_title = ?"
	exp.param = append(exp.param, MetaTitle)
	return exp
}

func (exp *DBPostORMExample) MetaTitleLike(MetaTitle string) *DBPostORMExample {
	exp.where += "meta_title like '%' || ? || '%'"
	exp.param = append(exp.param, MetaTitle)
	return exp
}

func (exp *DBPostORMExample) MetaTitleLikeAfter(MetaTitle string) *DBPostORMExample {
	exp.where += "meta_title like ? || '%'"
	exp.param = append(exp.param, MetaTitle)
	return exp
}

func (exp *DBPostORMExample) MetaTitleLikeBefore(MetaTitle string) *DBPostORMExample {
	exp.where += "meta_title like  '%' || ?"
	exp.param = append(exp.param, MetaTitle)
	return exp
}





func (exp *DBPostORMExample) MetaDescriptionEq(MetaDescription string) *DBPostORMExample {
	exp.where += "meta_description = ?"
	exp.param = append(exp.param, MetaDescription)
	return exp
}

func (exp *DBPostORMExample) MetaDescriptionLike(MetaDescription string) *DBPostORMExample {
	exp.where += "meta_description like '%' || ? || '%'"
	exp.param = append(exp.param, MetaDescription)
	return exp
}

func (exp *DBPostORMExample) MetaDescriptionLikeAfter(MetaDescription string) *DBPostORMExample {
	exp.where += "meta_description like ? || '%'"
	exp.param = append(exp.param, MetaDescription)
	return exp
}

func (exp *DBPostORMExample) MetaDescriptionLikeBefore(MetaDescription string) *DBPostORMExample {
	exp.where += "meta_description like  '%' || ?"
	exp.param = append(exp.param, MetaDescription)
	return exp
}




func (exp *DBPostORMExample) CreatedEq(Created time.Time) *DBPostORMExample {
	exp.where += "created = ?"
	exp.param = append(exp.param, Created)
	return exp
}
func (exp *DBPostORMExample) CreatedGt(Created time.Time) *DBPostORMExample {
	exp.where += "created > ?"
	exp.param = append(exp.param, Created)
	return exp
}
func (exp *DBPostORMExample) CreatedLt(Created time.Time) *DBPostORMExample {
	exp.where += "created < ?"
	exp.param = append(exp.param, Created)
	return exp
}
func (exp *DBPostORMExample) CreatedEqGt(Created time.Time) *DBPostORMExample {
	exp.where += "created >= ?"
	exp.param = append(exp.param, Created)
	return exp
}
func (exp *DBPostORMExample) CreatedEqLt(Created time.Time) *DBPostORMExample {
	exp.where += "created <= ?"
	exp.param = append(exp.param, Created)
	return exp
}





func (exp *DBPostORMExample) ModifiedEq(Modified time.Time) *DBPostORMExample {
	exp.where += "modified = ?"
	exp.param = append(exp.param, Modified)
	return exp
}
func (exp *DBPostORMExample) ModifiedGt(Modified time.Time) *DBPostORMExample {
	exp.where += "modified > ?"
	exp.param = append(exp.param, Modified)
	return exp
}
func (exp *DBPostORMExample) ModifiedLt(Modified time.Time) *DBPostORMExample {
	exp.where += "modified < ?"
	exp.param = append(exp.param, Modified)
	return exp
}
func (exp *DBPostORMExample) ModifiedEqGt(Modified time.Time) *DBPostORMExample {
	exp.where += "modified >= ?"
	exp.param = append(exp.param, Modified)
	return exp
}
func (exp *DBPostORMExample) ModifiedEqLt(Modified time.Time) *DBPostORMExample {
	exp.where += "modified <= ?"
	exp.param = append(exp.param, Modified)
	return exp
}





func (exp *DBPostORMExample) PublishTimeEq(PublishTime time.Time) *DBPostORMExample {
	exp.where += "publish_time = ?"
	exp.param = append(exp.param, PublishTime)
	return exp
}
func (exp *DBPostORMExample) PublishTimeGt(PublishTime time.Time) *DBPostORMExample {
	exp.where += "publish_time > ?"
	exp.param = append(exp.param, PublishTime)
	return exp
}
func (exp *DBPostORMExample) PublishTimeLt(PublishTime time.Time) *DBPostORMExample {
	exp.where += "publish_time < ?"
	exp.param = append(exp.param, PublishTime)
	return exp
}
func (exp *DBPostORMExample) PublishTimeEqGt(PublishTime time.Time) *DBPostORMExample {
	exp.where += "publish_time >= ?"
	exp.param = append(exp.param, PublishTime)
	return exp
}
func (exp *DBPostORMExample) PublishTimeEqLt(PublishTime time.Time) *DBPostORMExample {
	exp.where += "publish_time <= ?"
	exp.param = append(exp.param, PublishTime)
	return exp
}




func (exp *DBPostORMExample) StatusEq(Status int) *DBPostORMExample {
	exp.where += "status = ?"
	exp.param = append(exp.param, Status)
	return exp
}
func (exp *DBPostORMExample) StatusGt(Status int) *DBPostORMExample {
	exp.where += "status > ?"
	exp.param = append(exp.param, Status)
	return exp
}
func (exp *DBPostORMExample) StatusLt(Status int) *DBPostORMExample {
	exp.where += "status < ?"
	exp.param = append(exp.param, Status)
	return exp
}
func (exp *DBPostORMExample) StatusEqGt(Status int) *DBPostORMExample {
	exp.where += "status >= ?"
	exp.param = append(exp.param, Status)
	return exp
}
func (exp *DBPostORMExample) StatusEqLt(Status int) *DBPostORMExample {
	exp.where += "status <= ?"
	exp.param = append(exp.param, Status)
	return exp
}





func (exp *DBPostORMExample) OrderByIDDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id desc"
    return exp
}
func (exp *DBPostORMExample) OrderByIDAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id asc"
    return exp
}

func (exp *DBPostORMExample) OrderByTitleDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " title desc"
    return exp
}
func (exp *DBPostORMExample) OrderByTitleAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " title asc"
    return exp
}

func (exp *DBPostORMExample) OrderByContentDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " content desc"
    return exp
}
func (exp *DBPostORMExample) OrderByContentAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " content asc"
    return exp
}

func (exp *DBPostORMExample) OrderByContentHTMLDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " content_html desc"
    return exp
}
func (exp *DBPostORMExample) OrderByContentHTMLAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " content_html asc"
    return exp
}

func (exp *DBPostORMExample) OrderByExcerptDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " excerpt desc"
    return exp
}
func (exp *DBPostORMExample) OrderByExcerptAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " excerpt asc"
    return exp
}

func (exp *DBPostORMExample) OrderByImageDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " image desc"
    return exp
}
func (exp *DBPostORMExample) OrderByImageAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " image asc"
    return exp
}

func (exp *DBPostORMExample) OrderByURLDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " url desc"
    return exp
}
func (exp *DBPostORMExample) OrderByURLAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " url asc"
    return exp
}

func (exp *DBPostORMExample) OrderByChannelIDDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " channel_id desc"
    return exp
}
func (exp *DBPostORMExample) OrderByChannelIDAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " channel_id asc"
    return exp
}

func (exp *DBPostORMExample) OrderByTagsIDDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " tags_id desc"
    return exp
}
func (exp *DBPostORMExample) OrderByTagsIDAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " tags_id asc"
    return exp
}

func (exp *DBPostORMExample) OrderByVisibleDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " visible desc"
    return exp
}
func (exp *DBPostORMExample) OrderByVisibleAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " visible asc"
    return exp
}

func (exp *DBPostORMExample) OrderByAuthorDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " author desc"
    return exp
}
func (exp *DBPostORMExample) OrderByAuthorAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " author asc"
    return exp
}

func (exp *DBPostORMExample) OrderByMetaTitleDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_title desc"
    return exp
}
func (exp *DBPostORMExample) OrderByMetaTitleAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_title asc"
    return exp
}

func (exp *DBPostORMExample) OrderByMetaDescriptionDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_description desc"
    return exp
}
func (exp *DBPostORMExample) OrderByMetaDescriptionAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_description asc"
    return exp
}

func (exp *DBPostORMExample) OrderByCreatedDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " created desc"
    return exp
}
func (exp *DBPostORMExample) OrderByCreatedAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " created asc"
    return exp
}

func (exp *DBPostORMExample) OrderByModifiedDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " modified desc"
    return exp
}
func (exp *DBPostORMExample) OrderByModifiedAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " modified asc"
    return exp
}

func (exp *DBPostORMExample) OrderByPublishTimeDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " publish_time desc"
    return exp
}
func (exp *DBPostORMExample) OrderByPublishTimeAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " publish_time asc"
    return exp
}

func (exp *DBPostORMExample) OrderByStatusDesc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " status desc"
    return exp
}
func (exp *DBPostORMExample) OrderByStatusAsc() *DBPostORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " status asc"
    return exp
}


type DBPostUpdater struct {
    setSQL  string
    sql     string
    param   []interface{}
    showSQL bool
    db      *sqlx.DB
}


func (updater *DBPostUpdater) SetID(ID int) *DBPostUpdater {
	updater.setSQL += ", id = ?"
	updater.param = append(updater.param, ID)
	return updater
}

func (updater *DBPostUpdater) SetTitle(Title string) *DBPostUpdater {
	updater.setSQL += ", title = ?"
	updater.param = append(updater.param, Title)
	return updater
}

func (updater *DBPostUpdater) SetContent(Content string) *DBPostUpdater {
	updater.setSQL += ", content = ?"
	updater.param = append(updater.param, Content)
	return updater
}

func (updater *DBPostUpdater) SetContentHTML(ContentHTML string) *DBPostUpdater {
	updater.setSQL += ", content_html = ?"
	updater.param = append(updater.param, ContentHTML)
	return updater
}

func (updater *DBPostUpdater) SetExcerpt(Excerpt string) *DBPostUpdater {
	updater.setSQL += ", excerpt = ?"
	updater.param = append(updater.param, Excerpt)
	return updater
}

func (updater *DBPostUpdater) SetImage(Image string) *DBPostUpdater {
	updater.setSQL += ", image = ?"
	updater.param = append(updater.param, Image)
	return updater
}

func (updater *DBPostUpdater) SetURL(URL string) *DBPostUpdater {
	updater.setSQL += ", url = ?"
	updater.param = append(updater.param, URL)
	return updater
}

func (updater *DBPostUpdater) SetChannelID(ChannelID int) *DBPostUpdater {
	updater.setSQL += ", channel_id = ?"
	updater.param = append(updater.param, ChannelID)
	return updater
}

func (updater *DBPostUpdater) SetTagsID(TagsID string) *DBPostUpdater {
	updater.setSQL += ", tags_id = ?"
	updater.param = append(updater.param, TagsID)
	return updater
}

func (updater *DBPostUpdater) SetVisible(Visible int) *DBPostUpdater {
	updater.setSQL += ", visible = ?"
	updater.param = append(updater.param, Visible)
	return updater
}

func (updater *DBPostUpdater) SetAuthor(Author int) *DBPostUpdater {
	updater.setSQL += ", author = ?"
	updater.param = append(updater.param, Author)
	return updater
}

func (updater *DBPostUpdater) SetMetaTitle(MetaTitle string) *DBPostUpdater {
	updater.setSQL += ", meta_title = ?"
	updater.param = append(updater.param, MetaTitle)
	return updater
}

func (updater *DBPostUpdater) SetMetaDescription(MetaDescription string) *DBPostUpdater {
	updater.setSQL += ", meta_description = ?"
	updater.param = append(updater.param, MetaDescription)
	return updater
}

func (updater *DBPostUpdater) SetCreated(Created *time.Time) *DBPostUpdater {
	updater.setSQL += ", created = ?"
	updater.param = append(updater.param, Created)
	return updater
}

func (updater *DBPostUpdater) SetModified(Modified *time.Time) *DBPostUpdater {
	updater.setSQL += ", modified = ?"
	updater.param = append(updater.param, Modified)
	return updater
}

func (updater *DBPostUpdater) SetPublishTime(PublishTime *time.Time) *DBPostUpdater {
	updater.setSQL += ", publish_time = ?"
	updater.param = append(updater.param, PublishTime)
	return updater
}

func (updater *DBPostUpdater) SetStatus(Status int) *DBPostUpdater {
	updater.setSQL += ", status = ?"
	updater.param = append(updater.param, Status)
	return updater
}

func (updater *DBPostUpdater) Done() *DBPostORMExample {
    example := &DBPostORMExample{sql: updater.sql + updater.setSQL[1:], db: updater.db, showSQL: updater.showSQL}
    for _, p := range updater.param {
        example.AddParam(p)
    }
	return example
}

type DBPostModel struct {
	ID int
	Title string
	Content string
	ContentHTML string
	Excerpt string
	Image string
	URL string
	ChannelID int
	TagsID string
	Visible int
	Author int
	MetaTitle string
	MetaDescription string
	Created *time.Time
	Modified *time.Time
	PublishTime *time.Time
	Status int
	
}

type DBPostORM struct {
	ShowSQL bool
	DB      *sqlx.DB
	sql     string
}

func (orm *DBPostORM) Insert(model *DBPostModel) *DBPostORMExample {
	orm.sql = "INSERT INTO t_post"
	orm.sql += "(title,content,content_html,excerpt,image,url,channel_id,tags_id,visible,author,meta_title,meta_description,created,modified,publish_time,status)"
	orm.sql += "VALUES"
	orm.sql += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	example := &DBPostORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
	
	
	
	
	example.AddParam(model.Title)
	
	
	
	example.AddParam(model.Content)
	
	
	
	example.AddParam(model.ContentHTML)
	
	
	
	example.AddParam(model.Excerpt)
	
	
	
	example.AddParam(model.Image)
	
	
	
	example.AddParam(model.URL)
	
	
	
	example.AddParam(model.ChannelID)
	
	
	
	example.AddParam(model.TagsID)
	
	
	
	example.AddParam(model.Visible)
	
	
	
	example.AddParam(model.Author)
	
	
	
	example.AddParam(model.MetaTitle)
	
	
	
	example.AddParam(model.MetaDescription)
	
	
	
	example.AddParam(model.Created)
	
	
	
	example.AddParam(model.Modified)
	
	
	
	example.AddParam(model.PublishTime)
	
	
	
	example.AddParam(model.Status)
	
	
	return example
}
func (orm *DBPostORM) Update(model *DBPostModel) *DBPostORMExample {
    orm.sql = "UPDATE t_post SET "
    orm.sql += "title = ?,content = ?,content_html = ?,excerpt = ?,image = ?,url = ?,channel_id = ?,tags_id = ?,visible = ?,author = ?,meta_title = ?,meta_description = ?,created = ?,modified = ?,publish_time = ?,status = ?"
    example := &DBPostORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
    
    
    
    
    example.AddParam(model.Title)
    
    
    
    example.AddParam(model.Content)
    
    
    
    example.AddParam(model.ContentHTML)
    
    
    
    example.AddParam(model.Excerpt)
    
    
    
    example.AddParam(model.Image)
    
    
    
    example.AddParam(model.URL)
    
    
    
    example.AddParam(model.ChannelID)
    
    
    
    example.AddParam(model.TagsID)
    
    
    
    example.AddParam(model.Visible)
    
    
    
    example.AddParam(model.Author)
    
    
    
    example.AddParam(model.MetaTitle)
    
    
    
    example.AddParam(model.MetaDescription)
    
    
    
    example.AddParam(model.Created)
    
    
    
    example.AddParam(model.Modified)
    
    
    
    example.AddParam(model.PublishTime)
    
    
    
    example.AddParam(model.Status)
    
    
    return example
}
func (orm *DBPostORM) UpdateField() *DBPostUpdater {
	orm.sql = "UPDATE t_post SET "
	return &DBPostUpdater{
		sql:     orm.sql,
		param:   nil,
		showSQL: orm.ShowSQL,
		db:      orm.DB,
	}
}
func (orm *DBPostORM) Delete() *DBPostORMExample {
    orm.sql = "DELETE FROM t_post"
    return &DBPostORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}
func (orm *DBPostORM) Query() *DBPostORMExample {
    orm.sql = "SELECT id,title,content,content_html,excerpt,image,url,channel_id,tags_id,visible,author,meta_title,meta_description,created,modified,publish_time,status FROM t_post"
    return &DBPostORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}



func (orm *DBPostORM) QueryNoLarge() *DBPostORMExample {
    orm.sql = "SELECT status,excerpt,url,tags_id,visible,modified,created,publish_time,id,title,image,channel_id,author FROM t_post"
    return &DBPostORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}

