package core

import (
	"encoding/json"
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
	err = cfg.replaceValue(cfg.cfgMap)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) replaceValue(cfgMap map[string]interface{}) error {
	for key, val := range cfgMap {
		vt := reflect.TypeOf(val)
		fmt.Println(key, vt.Kind())
		if vt.Kind().String() == "map" {
			cfg.replaceValue(val.(map[string]interface{}))
			continue
		}
		if vt.Kind().String() == "string" {
			has, reps := cfg.hasReplace(val.(string))
			if !has {
				continue
			}
			for _, str := range reps {
				
			}
		}
	}
	return nil
}

//func (cfg *Config) replaceValueByKey(key string) error {
//value, exists := cfg.Get(key)
//if !exists {
//	return errors.New(fmt.Sprintf("%v is not exists", key))
//}
//if hasReplace, repKey := cfg.hasReplace(value); hasReplace {
//	data, ok := cfg.Get(repKey)
//	if !ok {
//		return errors.New(fmt.Sprintf("%v is not exists", repKey))
//	}
//	if hr, _ := cfg.hasReplace(data.(string)); hr {
//		cfg.replaceValueByKey(repKey, cfgMap)
//	}
//	cfgMap[key] = hasReplaceRegex.ReplaceAllString(cfgMap[key].(string), data.(string))
//}
//return nil
//}

func (cfg *Config) hasReplace(value string) (has bool, key []string) {
	vals := hasReplaceRegex.FindAllStringSubmatch(value, -1)
	if len(vals) == 0 {
		return false, []string{}
	}
	key = vals[0]
	has = len(key) > 0
	return
}

//func (cfg *Config) Get(key string) (value interface{}, exists bool, dataType string) {
//	keys := strings.Split(key, ".")
//	getValue := func(tmpMap map[string]interface{}) (value interface{}, exists bool, dataType string) {
//		length := len(keys) - 1
//		for i, k := range keys {
//			val := tmpMap[k]
//			tv := reflect.TypeOf(val)
//			if length == i {
//				return val, true, tv.String()
//			}
//		}
//		return nil, false, ""
//	}
//	return getValue(cfg.cfgMap)
//}
