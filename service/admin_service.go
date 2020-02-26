package service

import (
	"crypto/md5"
	"encoding/hex"
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

func (s *AdminService) GenToken(admin *model.Admin) {
	uid := uuid.Must(uuid.NewV4())
	token := strings.ReplaceAll(uid.String(), "-", "")
	cacheKey := "token:" + token
	admin.Token = token
	s.CacheService.Set(cacheKey, admin, 30*time.Minute)
}

func (s *AdminService) CheckToken(token string) (*model.Admin, error) {
	cacheKey := "token:" + token
	admin := &model.Admin{}
	if err := s.CacheService.Get(cacheKey, admin); err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *AdminService) TrieGet(id uint) *model.Admin {
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
