package test

import (
	"encoding/json"
	"fmt"
	"github.com/wang22/easyjson"
	"io/ioutil"
	"testing"
)

func TestYaml(t *testing.T) {
	f, _ := ioutil.ReadFile("../app_config/config_dev.json")
	var data map[string]interface{}
	json.Unmarshal(f, &data)
	fmt.Println(data)
	pj := easyjson.ParseJSON(f)
	fmt.Println(pj.GetObject("cache").GetInt("db_index"))
}

func TestJson(t *testing.T) {
	f, _ := ioutil.ReadFile("../app_config/config_dev.json")
	easy := easyjson.ParseJSON(f)
	//key := "cache.type"
	fmt.Println(easy.ChainCall("cache.type"))

}
