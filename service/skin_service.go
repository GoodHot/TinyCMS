package service

import (
	"encoding/json"
	"github.com/GoodHot/TinyCMS/common/files"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

type SkinService struct {
	ORM       *orm.ORM `ioc:"auto"`
	SkinDir   string   `val:"${server.skin_template_dir}"`
	StaticDir string   `val:"${server.static}"`
}

type SkinInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Template    []struct {
		File        string `json:"file"`
		Description string `json:"description"`
	}
	CoverImgs map[string]string `json:"cover_imgs"`
}

func (s *SkinService) AllSkin() ([]*SkinInfo, error) {
	dirs, err := ioutil.ReadDir(s.SkinDir)
	if err != nil {
		return nil, err
	}
	var skins []*SkinInfo
	for _, dir := range dirs {
		infoPath := s.SkinDir + "/" + dir.Name() + "/info.json"
		infoJSON, err := ioutil.ReadFile(infoPath)
		if err != nil {
			continue
		}
		skin := &SkinInfo{}
		err = json.Unmarshal(infoJSON, skin)
		if err != nil {
			log.Error(err)
			continue
		}
		coverPath := s.SkinDir + "/" + dir.Name() + "/" + skin.Cover
		imgs, err := ioutil.ReadDir(coverPath)
		if err != nil {
			log.Error(err)
		} else {
			coverImgs := make(map[string]string)
			s.checkSkinCover2Static(dir.Name(), coverPath)
			for _, img := range imgs {
				coverImgs[img.Name()] = "/" + s.StaticDir + "/skin_cover/" + dir.Name() + "/" + img.Name()
			}
			skin.CoverImgs = coverImgs
		}
		skins = append(skins, skin)
	}
	return skins, nil
}

func (s *SkinService) checkSkinCover2Static(name, coverPath string) bool {
	// TODO 这里存在BUG，当模板更新后，并不会同步过去，以后修复
	dir := s.StaticDir + "/skin_cover/" + name
	if (files.DirExists(dir)) {
		return true
	}
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return false
	}
	filepath.Walk(coverPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files.CopyFile(dir + "/" + info.Name(), path)
		return nil
	})
	return true
}
