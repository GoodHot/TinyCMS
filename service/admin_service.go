package service

import (
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type AdminService struct {
	AdminORM *orm.AdminORMFactory `ioc:"auto"`
}

func (*AdminService) GetByID(id int) {

}

func (s *AdminService) GetByUsernameOrEmail(account string) (*trait.Admin, error) {
	return s.AdminORM.Get().GetByUsername(account)
}

func (*AdminService) CheckPwd(oldPwd, newPwd string) bool {
	return oldPwd == newPwd
}
