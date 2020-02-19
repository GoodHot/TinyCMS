package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type AdminService struct {
	ORM *orm.ORM `ioc:"auto"`
}

func (s *AdminService) GetByUsername(username string) *model.Admin {
	model := &model.Admin{}
	err := s.ORM.DB.Where("username = ?", username).Find(model).Error
	if err != nil {
		return nil
	}
	return model
}

func (s *AdminService) AssertPwd(pwd string, enPwd string) bool {
	pwd = s.encrptyPwd(pwd)
	return pwd == enPwd
}

func (t *AdminService) encrptyPwd(pwd string) string {
	h1 := md5.New()
	h1.Write([]byte(pwd))
	pwdOnce := hex.EncodeToString(h1.Sum(nil))
	h2 := md5.New()
	h2.Write([]byte(pwdOnce[16:] + pwdOnce[:16]))
	return hex.EncodeToString(h2.Sum(nil))
}
