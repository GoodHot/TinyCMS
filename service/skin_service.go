package service

import (
	"encoding/json"
	"errors"
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/consts"
	"github.com/GoodHot/TinyCMS/common/files"
	"github.com/GoodHot/TinyCMS/common/render"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"time"
)

type SkinService struct {
	Log                brain.Logger   `ioc:"auto"`
	ORM                *orm.ORM       `ioc:"auto"`
	ConfigService      *ConfigService `ioc:"auto"`
	CacheService       *CacheService  `ioc:"auto"`
	SkinTemplateDir    string         `val:"${server.skin_template_dir}"`
	ServerTemplateDir  string         `val:"${server.skin_template_dir}"` // 模板文件夹
	ServerHTMLCompress bool           `val:"${server.html_compress}"`     // HTML代码是否压缩
	htmlRender         *render.HTMLRenderer
}

type SkinInfo struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Version     string         `json:"version"`
	Author      string         `json:"author"`
	Website     string         `json:"website"`
	Pages       []SkinInfoPage `json:"pages"`
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
	s.CacheService.Delete("skin:active")
	var existsSkins []string
	// 逐个加载模板文件夹
	for _, dir := range dirs {
		s.Log.Info("Find skin template directory [%v]", dir.Name())
		skinDir := strs.Fmt("%s/%s", s.SkinTemplateDir, dir.Name())
		err := s.loadSkin(dir.Name(), skinDir)
		if err != nil {
			s.Log.Error("Load skin template [%s] failed, error message : %s", skinDir, err.Error())
			return err
		}
		existsSkins = append(existsSkins, dir.Name())
	}
	s.ORM.DB.Unscoped().Delete(&model.Skin{}, "dir not in (?)", existsSkins)
	// 判断是否有默认模板，如果没有，需要指定一个
	active := s.GetActive()
	if active != nil {
		return nil
	}
	skin := &model.Skin{}
	s.ORM.DB.First(skin)
	return s.ORM.DB.Model(skin).Where("id = ?", skin.ID).Update("active", true).Error
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
	skin := &model.Skin{}
	cacheKey := "skin:active"
	s.CacheService.Get(cacheKey, skin)
	if skin.Dir == "" {
		s.ORM.DB.Where("active = ?", true).First(skin)
		if skin.Dir == "" {
			return nil
		}
		s.CacheService.Set(cacheKey, skin, 24*time.Hour)
	}
	return skin
}

func (s *SkinService) List() {
}

func (s *SkinService) initHTMLRender() error {
	skin := s.GetActive()
	s.htmlRender = &render.HTMLRenderer{
		LayoutDir:    "layout",
		ComponentDir: "component",
		TemplateDir:  strs.Fmt("%s/%s", s.ServerTemplateDir, skin.Dir),
		Suffix:       ".html",
		Compress:     s.ServerHTMLCompress,
	}
	return s.htmlRender.Init(nil)
}

func (s *SkinService) SwitchSkin(id int) error {
	skin, err := s.GetById(id)
	if err != nil {
		return err
	}
	active := s.GetActive()
	err = s.ORM.DB.Model(&model.Skin{}).Where("id = ?", active.ID).Update("active", false).Error
	if err != nil {
		return err
	}
	err = s.ORM.DB.Model(&model.Skin{}).Where("id = ?", skin.ID).Update("active", true).Error
	if err != nil {
		return err
	}
	return s.CacheService.Set("skin:active", skin, 24*time.Hour)
}

func (s *SkinService) GetById(id int) (*model.Skin, error) {
	var skin model.Skin
	s.ORM.DB.Where("id = ?", id).First(&skin)
	if skin.Dir == "" {
		return nil, errors.New(strs.Fmt("Skin is %v not exists", id))
	}
	return &skin, nil
}

func (s *SkinService) Render(writer io.Writer, name string, data interface{}, ctx echo.Context) error {
	if s.htmlRender == nil {
		err := s.initHTMLRender()
		if err != nil {
			s.Log.Error("Render Skin template failed, error message is :", err.Error())
			return err
		}
	} else {
		// TODO 这里可能会有性能问题，但这样可以处理分布式部署问题
		active := s.GetActive()
		if s.htmlRender.TemplateDir != strs.Fmt("%s/%s", s.ServerTemplateDir, active.Dir) {
			err := s.initHTMLRender()
			if err != nil {
				s.Log.Error("Render Skin template failed, error message is :", err.Error())
				return err
			}
		}
	}
	return s.htmlRender.Render(writer, name, data)
}
