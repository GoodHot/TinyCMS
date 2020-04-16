package service

import (
	"encoding/json"
	"errors"
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/consts"
	"github.com/GoodHot/TinyCMS/common/files"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"io/ioutil"
	"time"
)

type SkinService struct {
	Log             brain.Logger   `ioc:"auto"`
	ORM             *orm.ORM       `ioc:"auto"`
	ConfigService   *ConfigService `ioc:"auto"`
	SkinTemplateDir string         `val:"${server.skin_template_dir}"`
}

type SkinInfo struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Version     string         `json:"version"`
	Author      string         `json:"author"`
	Website     string         `json:"website"`
	Pages       []SkinInfoPage `json:"pages"`
	Cover       []struct {
		Image       string `json:"image"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"cover"`
}

type SkinInfoPage struct {
	Page        string `json:"page"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *SkinService) Startup() error {
	s.Log.Info("Initial skin template from [%s]", s.SkinTemplateDir)
	// 读取模板文件夹，每个文件夹代表一套模板
	dirs, err := ioutil.ReadDir(s.SkinTemplateDir)
	if err != nil {
		return err
	}
	if dirs == nil || len(dirs) == 0 {
		s.Log.Warn("Skin template [%s] no files found", s.SkinTemplateDir)
		return nil
	}
	// 逐个加载模板文件夹
	for _, dir := range dirs {
		s.Log.Info("Find skin template directory [%v]", dir.Name())
		skinDir := strs.Fmt("%s/%s", s.SkinTemplateDir, dir.Name())
		err := s.loadSkin(dir.Name(), skinDir)
		if err != nil {
			s.Log.Error("Load skin template [%s] failed, error message : %s", skinDir, err.Error())
			return err
		}
	}
	// 判断是否有默认模板，如果没有，需要指定一个
	active := s.GetActive()
	if active != nil {
		return nil
	}
	skin := &model.Skin{}
	s.ORM.DB.First(skin)
	s.ORM.DB.Model(skin).Where("id = ?", skin.ID).Update("active", true)
	return nil
}

func (s *SkinService) loadSkin(dirName string, skinDir string) error {
	// 读取skin.json文件
	jsonPath := strs.Fmt("%s/%s", skinDir, consts.SkinTemplateJsonName())
	skinJSON, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	skinInfo := &SkinInfo{}
	err = json.Unmarshal(skinJSON, skinInfo)
	if err != nil {
		return err
	}
	// 获取模板页面数据
	pageMap := make(map[string]SkinInfoPage)
	for _, page := range skinInfo.Pages {
		filePath := strs.Fmt("%s/%s", skinDir, page.Page)
		if !files.FileExists(filePath) {
			return errors.New(strs.Fmt("Skin template can not find page file %s", filePath))
		}
		pageMap[page.Page] = page
	}
	// 检查是否有缺失文件
	var lostFiles []string
	for _, page := range consts.SkinTemplateContainPage() {
		if _, ok := pageMap[page]; !ok {
			lostFiles = append(lostFiles, page)
		}
	}
	if lostFiles != nil && len(lostFiles) != 0 {
		msg := strs.Fmt("Skin template must include %v", lostFiles)
		s.Log.Error(msg)
		return errors.New(msg)
	}
	activeSkin := s.GetActive()
	active := activeSkin == nil && dirName == "default"
	s.Save(skinInfo, dirName, active)
	return nil
}

func (s *SkinService) Save(skinInfo *SkinInfo, skinDir string, active bool) {
	skin := &model.Skin{}
	s.ORM.DB.Where("dir = ?", skinDir).First(skin)
	skin.Dir = skinDir
	skin.Name = skinInfo.Name
	skin.Description = skinInfo.Description
	skin.Version = skinInfo.Version
	skin.Author = skinInfo.Author
	skin.Website = skinInfo.Website
	if active {
		skin.Active = active
	}
	skin.CreatedAt = time.Time{}
	skin.UpdatedAt = time.Time{}
	s.ORM.DB.Save(skin)
}

func (s *SkinService) GetActive() *model.Skin {
	var skin model.Skin
	s.ORM.DB.Where("active = ?", 1).First(&skin)
	if skin.Dir == "" {
		return nil
	}
	return &skin
}

func (s *SkinService) List() {
}

/**
安装皮肤
*/
func (s *SkinService) Install() {
}
