package service

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"strings"
)

type MemberService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (*MemberService) GetByID(id int) {

}

func (s *MemberService) GetByUsernameOrEmail(account string) (*trait.Member, *core.Err) {
	if strings.TrimSpace(account) == "" {
		return nil, core.NewErr(core.Err_Auth_Account_Fail)
	}
	if strings.Contains(account, "@") {
		return s.ORM.Member.GetByEmail(account)
	}
	return s.ORM.Member.GetByUsername(account)
}

func (*MemberService) CheckPwd(oldPwd, newPwd string) bool {
	return oldPwd == newPwd
}

func (s *MemberService) Page(query *trait.Query) (*trait.Page, *core.Err) {
	return s.ORM.Member.Page(query)
}
