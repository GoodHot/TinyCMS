package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type ConfigService struct {
	ORM *orm.ORM `ioc:"auto"`
}

func (s *ConfigService) InitConfig(cfg *model.Config) {
	old := &model.Config{}
	s.ORM.DB.Where("key = ?", cfg.Key).First(old)
	if old.Key == "" {
		s.ORM.DB.Save(cfg)
	}
}
