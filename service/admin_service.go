package service

import (
	"errors"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"strings"
)

type AdminService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (*AdminService) GetByID(id int) {

}

func (s *AdminService) GetByUsernameOrEmail(account string) (*trait.Admin, error) {
	if strings.TrimSpace(account) == "" {
		return nil, errors.New("not found")
	}
	if strings.Contains(account, "@") {
		return s.ORM.Admin.GetByEmail(account)
	}
	return s.ORM.Admin.GetByUsername(account)
}

func (*AdminService) CheckPwd(oldPwd, newPwd string) bool {
	return oldPwd == newPwd
}
