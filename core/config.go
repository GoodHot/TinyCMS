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
	cfg.replaceValue(cfg.cfgMap)
	fmt.Println(cfg.cfgMap)
	return nil
}

func (cfg *Config) replaceValue(cfgMap map[string]interface{}) error {
	for key, val := range cfgMap {
		vt := reflect.TypeOf(val)
		if vt.Kind().String() == "string" {
			cfg.replaceValueByKey(key, cfgMap)
		}
	}
	return nil
}

func (cfg *Config) replaceValueByKey(key string, cfgMap map[string]interface{}) error {
	value := cfgMap[key]
	if hasReplace, repKey := cfg.hasReplace(value.(string)); hasReplace {
		data, ok := cfgMap[repKey]
		if !ok {
			return errors.New(fmt.Sprintf("%v is not exists", repKey))
		}
		if hr, _ := cfg.hasReplace(data.(string)); hr {
			cfg.replaceValueByKey(repKey, cfgMap)
		}
		cfgMap[key] = hasReplaceRegex.ReplaceAllString(cfgMap[key].(string), data.(string))
	}
	return nil
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
