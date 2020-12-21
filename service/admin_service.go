package service

import (
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type AdminService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (*AdminService) GetByID(id int) {

}

func (s *AdminService) GetByUsernameOrEmail(account string) (*trait.Admin, error) {
	return s.ORM.Admin.GetByUsername(account)
}

func (*AdminService) CheckPwd(oldPwd, newPwd string) bool {
	return oldPwd == newPwd
}
