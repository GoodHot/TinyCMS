package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
)

var (
	hasReplaceRegex, _ = regexp.Compile(`\$\{(.+?)\}`)
)

type Config struct {
	cfgMap map[string]interface{}
}

func (cfg *Config) ReadConfig(cfgPath string) error {
	cfgData, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}
	cfg.cfgMap = make(map[string]interface{})
	err = json.Unmarshal(cfgData, &cfg.cfgMap)
	if err != nil {
		return err
	}
	cfg.replaceValue()
	return nil
}

func (cfg *Config) replaceValue() error {
	for k, v := range cfg.cfgMap {
		fmt.Println(k, v)
		//vt := reflect.TypeOf(v)
		//if vt.Kind().String() == "string" {
		//	if has, repKey := cfg.hasReplace(v.(string)); has {
		//		fmt.Println(k, v, repKey)
		//		data, exists := cfg.Get(repKey)
		//		if !exists {
		//			return errors.New(fmt.Sprintf("%v is not exists"))
		//		}
		//
		//	}
		//}
	}
}

func (cfg *Config) hasReplace(value string) (has bool, key string) {
	vals := hasReplaceRegex.FindAllStringSubmatch(value, -1)
	if len(vals) == 0 {
		return false, ""
	}
	key = vals[0][1]
	has = key != ""
	return
}

func (*Config) Get(key string) (string, bool) {
	return "", false
}
