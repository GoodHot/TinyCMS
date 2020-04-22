package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/importcjj/trie-go"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

type AdminService struct {
	ORM          *orm.ORM      `ioc:"auto"`
	CacheService *CacheService `ioc:"auto"`
	dataTree     *trie.Trie
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
	pwd = s.EncrptyPwd(pwd)
	return pwd == enPwd
}

func (t *AdminService) EncrptyPwd(pwd string) string {
	h1 := md5.New()
	h1.Write([]byte(pwd))
	pwdOnce := hex.EncodeToString(h1.Sum(nil))
	h2 := md5.New()
	h2.Write([]byte(pwdOnce[16:] + pwdOnce[:16]))
	return hex.EncodeToString(h2.Sum(nil))
}

func (s *AdminService) GenToken(admin *model.Admin) *model.AdminResult {
	uid := uuid.Must(uuid.NewV4(), nil)
	token := strings.ReplaceAll(uid.String(), "-", "")
	cacheKey := "token:" + token
	s.CacheService.Set(cacheKey, admin, 60*time.Minute)
	result := &model.AdminResult{
		Admin:     admin,
		Token:     token,
		CreatedAt: nil,
		UpdatedAt: nil,
		DeletedAt: nil,
	}
	return result
}

func (s *AdminService) CheckToken(token string) (*model.AdminResult, error) {
	cacheKey := "token:" + token
	admin := &model.AdminResult{}
	if err := s.CacheService.Get(cacheKey, admin); err != nil {
		return nil, err
	}
	return admin, nil
}

//
//func (s *AdminService) CheckToken(token string) (*model.Admin, error) {
//	cacheKey := "token:" + token
//	admin := &model.Admin{}
//	if err := s.CacheService.Get(cacheKey, admin); err != nil {
//		return nil, err
//	}
//	return admin, nil
//}
//
func (s *AdminService) TrieGet(id int) *model.Admin {
	if s.dataTree == nil {
		s.dataTree = trie.New()
		page, size := 1, 20
		row := (page - 1) * size
		var admins []*model.Admin
		for ; ; {
			s.ORM.DB.Limit(size).Offset(row).Find(&admins)
			if len(admins) == 0 {
				break
			}
			for _, adm := range admins {
				s.dataTree.Put(strconv.Itoa(int(adm.ID)), adm)
			}
			page++
			row = (page - 1) * size
		}
	}
	exists, result := s.dataTree.Match(strconv.Itoa(int(id)))
	if !exists {
		return nil
	}
	return result.Value.(*model.Admin)
}

func (s *AdminService) All(selfID int) []*model.Admin {
	var result []*model.Admin
	s.ORM.DB.Find(&result)
	//if result != nil && len(result) > 0 {
		//for _, item := range result {
			//role, _ := s.RoleService.Get(int(item.RoleID))
			//item.Role = role
		//	item.IsSelf = item.ID == selfID
		//}
	//}
	return result
}

func (s *AdminService) Save(admin *model.Admin) error {
	var exists model.Admin
	s.ORM.DB.Where("username = ? and id != ?", admin.Username, admin.ID).First(&exists)
	if exists.ID != 0 {
		return errors.New(strs.Fmt("用户名'%s'已存在", admin.Username))
	}
	admin.Password = s.EncrptyPwd(admin.Password)
	return s.ORM.DB.Save(admin).Error
}

//
//func (s *AdminService) All(selfID uint) []*model.Admin {
//	var result []*model.Admin
//	s.ORM.DB.Find(&result)
//	if result != nil && len(result) > 0 {
//		for _, item := range result {
//			role, _ := s.RoleService.Get(int(item.RoleID))
//			item.Role = role
//			item.IsSelf = item.ID == selfID
//		}
//	}
//	return result
//}
//
//func (s *AdminService) Save(admin *model.Admin) error {
//	var exists model.Admin
//	s.ORM.DB.Where("username = ? and id != ?", admin.Username, admin.ID).First(&exists)
//	if exists.ID != 0 {
//		return errors.New(strs.Fmt("用户名'%s'已存在", admin.Username))
//	}
//	if admin.ID > 0 && admin.Password == "no_change_password!@#" {
//		adm := s.Get(int(admin.ID))
//		admin.Password = adm.Password
//	} else {
//		admin.Password = s.encrptyPwd(admin.Password)
//	}
//	return s.ORM.DB.Save(admin).Error
//}
//
//func (s *AdminService) Get(id int) *model.Admin {
//	var admin model.Admin
//	s.ORM.DB.Where("id = ?", id).Find(&admin)
//	return &admin
//}
//
//func (s *AdminService) Delete(id int) error {
//	count := 0
//	s.ORM.DB.Model(&model.Admin{}).Count(&count)
//	if count < 2 {
//		return errors.New("不允许删除最后一个用户")
//	}
//	return s.ORM.DB.Where("id = ?", id).Delete(&model.Admin{}).Error
//}
