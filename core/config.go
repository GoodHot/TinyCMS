package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
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
	for again := true; again; {
		again, err = cfg.replaceValue(cfg.cfgMap)
		if err != nil {
			return err
		}
	}
	fmt.Println(cfg.cfgMap)
	return nil
}

func (cfg *Config) replaceValue(cfgMap map[string]interface{}) (again bool, err error) {
	needAgain := false
	for key, val := range cfgMap {
		vt := reflect.TypeOf(val)
		if vt.Kind().String() == "map" {
			cfg.replaceValue(val.(map[string]interface{}))
			continue
		}
		if vt.Kind().String() == "string" {
			has, keys := cfg.hasReplace(val.(string))
			if !has {
				continue
			}
			tmpVal := val.(string)
			for _, k := range keys {
				str, err := cfg.GetStr(k)
				if err != nil {
					return false, err
				}
				tmpVal = strings.ReplaceAll(tmpVal, fmt.Sprintf("${%v}", k), str)
			}
			has, _ = cfg.hasReplace(tmpVal)
			needAgain = !needAgain && has
			cfgMap[key] = tmpVal
		}
	}
	return needAgain, nil
}

func (cfg *Config) GetStr(key string) (value string, err error) {
	keyChain := strings.Split(key, ".")
	swapMap := cfg.cfgMap
	length := len(keyChain) - 1
	for i, k := range keyChain {
		val := swapMap[k]
		vt := reflect.TypeOf(val).Kind().String()
		if i == length {
			if vt != "string" {
				return "", errors.New(fmt.Sprintf("%v is not string", key))
			}
			return val.(string), nil
		}
		if vt == "map" {
			swapMap = val.(map[string]interface{})
		} else {
			return "", errors.New(fmt.Sprintf("can not find key %v", key))
		}
	}
	return
}

func (cfg *Config) hasReplace(value string) (has bool, keys []string) {
	match := hasReplaceRegex.FindAllStringSubmatch(value, -1)
	if len(match) == 0 {
		return false, []string{}
	}
	has = len(match) > 0
	if !has {
		return has, nil
	}
	for _, val := range match {
		keys = append(keys, val[1])
	}
	return
}
