package test

import (
	"encoding/json"
	"fmt"
	"github.com/GoodHot/TinyCMS/service"
	"testing"
)

func TestSkinDir(t *testing.T) {
	service := service.SkinService{
		SkinDir:   "../resource/skin",
		StaticDir: "../static",
	}
	skins, err := service.AllSkin()
	rst, _ := json.Marshal(skins)
	fmt.Println(string(rst), err)
}
