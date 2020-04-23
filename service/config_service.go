package service

import (
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type ConfigService struct {
	ORM *orm.ORM     `ioc:"auto"`
	Log brain.Logger `ioc:"auto"`
}

func (s *ConfigService) Startup() error {
	s.Log.Info("Start initial system config")
	s.createConfig(&model.Config{
		Name:        "站点名称",
		Description: "站点名称",
		Key:         "website_name",
		Value:       "站点名称",
		OptType:     model.OptTypeText,
		Visible:     true,
		ConfigGroup: model.DictGroupSystem,
	})
	s.createConfig(&model.Config{
		Name:        "站点域名",
		Description: "站点域名",
		Key:         "website_domain",
		Value:       "false.run",
		OptType:     model.OptTypeText,
		Visible:     true,
		ConfigGroup: model.DictGroupSystem,
	})
	s.createConfig(&model.Config{
		Name:        "关键字",
		Description: "seo关键字",
		Key:         "seo_keyword",
		Value:       "tiny,cms",
		OptType:     model.OptTypeTags,
		Visible:     true,
		ConfigGroup: model.DictGroupSEO,
	})
	s.createConfig(&model.Config{
		Name:        "描述",
		Description: "seo描述",
		Key:         "seo_description",
		Value:       "tiny,cms",
		OptType:     model.OptTypeTextArea,
		Visible:     true,
		ConfigGroup: model.DictGroupSEO,
	})
	return nil
}

func (s *ConfigService) createConfig(config *model.Config) error {
	cfg := s.GetByKey(config.Key)
	if cfg.Name != "" {
		return nil
	}
	return s.ORM.DB.Save(config).Error
}

func (s *ConfigService) GetByKey(key string) *model.Config {
	cfg := &model.Config{}
	s.ORM.DB.Where("key = ?", key).First(cfg)
	return cfg
}

func (s *ConfigService) GetByGroup(group string) []*model.Config {
	var configs []*model.Config
	s.ORM.DB.Where("config_group = ?", group).Find(&configs)
	return configs
}

func (s *ConfigService) Update(configs []*model.Config) error {
	if configs == nil || len(configs) == 0 {
		return nil
	}
	for _, cfg := range configs {
		err := s.ORM.DB.Model(&model.Config{}).Where("key = ?", cfg.Key).Update("value", cfg.Value).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//func (s *ConfigService) InitConfig(cfg *model.Config) {
//	old := &model.Config{}
//	s.ORM.DB.Where("key = ?", cfg.Key).First(old)
//	if old.Key == "" {
//		s.ORM.DB.Save(cfg)
//	}
//}
