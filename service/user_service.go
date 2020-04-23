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

type UserService struct {
	ORM          *orm.ORM      `ioc:"auto"`
	CacheService *CacheService `ioc:"auto"`
	dataTree     *trie.Trie
}

func (s *UserService) GetByUsername(username string) *model.User {
	model := &model.User{}
	err := s.ORM.DB.Where("username = ?", username).Find(model).Error
	if err != nil {
		return nil
	}
	return model
}

func (s *UserService) AssertPwd(pwd string, enPwd string) bool {
	pwd = s.EncrptyPwd(pwd)
	return pwd == enPwd
}

func (s *UserService) EncrptyPwd(pwd string) string {
	h1 := md5.New()
	h1.Write([]byte(pwd))
	pwdOnce := hex.EncodeToString(h1.Sum(nil))
	h2 := md5.New()
	h2.Write([]byte(pwdOnce[16:] + pwdOnce[:16]))
	return hex.EncodeToString(h2.Sum(nil))
}

func (s *UserService) GenToken(user *model.User) *model.UserResult {
	uid := uuid.Must(uuid.NewV4(), nil)
	token := strings.ReplaceAll(uid.String(), "-", "")
	cacheKey := "token:" + token
	s.CacheService.Set(cacheKey, user, 60*time.Minute)
	result := &model.UserResult{
		User:      user,
		Token:     token,
		CreatedAt: nil,
		UpdatedAt: nil,
		DeletedAt: nil,
	}
	return result
}

func (s *UserService) CheckToken(token string) (*model.UserResult, error) {
	cacheKey := "token:" + token
	user := &model.UserResult{}
	if err := s.CacheService.Get(cacheKey, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) TrieGet(id int) *model.User {
	if s.dataTree == nil {
		s.dataTree = trie.New()
		page, size := 1, 20
		row := (page - 1) * size
		var users []*model.User
		for ; ; {
			s.ORM.DB.Limit(size).Offset(row).Find(&users)
			if len(users) == 0 {
				break
			}
			for _, adm := range users {
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
	return result.Value.(*model.User)
}

func (s *UserService) All(selfID int) []*model.User {
	var result []*model.User
	s.ORM.DB.Find(&result)
	return result
}

func (s *UserService) Save(user *model.User) error {
	var exists model.User
	s.ORM.DB.Where("username = ? and id != ?", user.Username, user.ID).First(&exists)
	if exists.ID != 0 {
		return errors.New(strs.Fmt("用户名'%s'已存在", user.Username))
	}
	user.Password = s.EncrptyPwd(user.Password)
	return s.ORM.DB.Save(user).Error
}
