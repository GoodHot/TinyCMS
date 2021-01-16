package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
	_ "github.com/mattn/go-sqlite3"
)

type MemberORMImpl struct {
	DB *datasource.DBMemberORM
}

func (orm *MemberORMImpl) Page(query *trait.Query) (*trait.Page, *core.Err) {
	exp := orm.DB.Query()
	if query.Param != nil {
		if query.Param["staff"] != nil && query.Param["staff"].(bool) {
			exp.RoleLt(int(trait.RoleTypeMember))
		} else {
			exp.RoleEqGt(int(trait.RoleTypeMember))
		}
		if query.Param["role"] != nil && query.Param["role"].(int) > 0 {
			exp.And().RoleEq(query.Param["role"].(int))
		}
	}
	count, err := exp.ExecQueryCount()
	if err != nil {
		return nil, core.NewErr(core.Err_Member_Get_Page_Fail)
	}
	exp.Limit(query.Size, (query.Page-1)*query.Size)
	if query.Order != nil {
		if query.Order["id"] == "desc" {
			exp.OrderByIDDesc()
		} else {
			exp.OrderByIDAsc()
		}
	}
	rst, err := exp.ExecQuery()
	if err != nil {
		return nil, core.NewErr(core.Err_Member_Get_Page_Fail)
	}
	page := &trait.Page{}
	page.Page = query.Page
	page.Size = query.Size
	total := count / page.Size
	if count%page.Size > 0 {
		total++
	}
	page.Count = count
	page.Total = total
	page.List = orm.convert(rst...)
	return page, nil
}

func (orm *MemberORMImpl) GetByUsername(username string) (*trait.Member, *core.Err) {
	model, err := orm.DB.Query().UsernameEq(username).ExecQueryOne()
	if err != nil || model == nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	return &orm.convert(model)[0], nil
}

func (orm *MemberORMImpl) GetByEmail(email string) (*trait.Member, *core.Err) {
	model, err := orm.DB.Query().EmailEq(email).ExecQueryOne()
	if err != nil || model == nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	return &orm.convert(model)[0], nil
}

func (orm *MemberORMImpl) convert(model ...*datasource.DBMemberModel) []trait.Member {
	var rst []trait.Member
	for _, item := range model {
		rst = append(rst, trait.Member{
			BaseORM: trait.BaseORM{
				ID: item.ID,
			},
			Nickname: item.Nickname,
			Username: item.Username,
			Email:    item.Email,
			Password: item.Password,
			Role:     trait.RoleType(item.Role),
		})
	}
	return rst
}
