package spring

import (
	"errors"
	"fmt"
	"github.com/wang22/easyjson"
	"reflect"
	"regexp"
)

type SpringIOC struct {
	instanceMap map[string]interface{}
	config      *easyjson.EasyJSON
}

func (ioc *SpringIOC) Reg(x interface{}) (ins interface{}, err error) {
	insType := reflect.TypeOf(x)
	if insType.Kind() != reflect.Ptr {
		return ins, errors.New("register must be struct")
	}
	name := ioc.instanceName(insType)
	ins = ioc.GetByName(name)
	if ins != nil {
		return
	}
	ioc.setIns(name, x)
	return ioc.createIns(x), nil
}

func (ioc *SpringIOC) Get(x interface{}) interface{} {
	refType := reflect.TypeOf(x)
	return ioc.GetByName(ioc.instanceName(refType))
}

func (ioc *SpringIOC) setIns(name string, x interface{}) error {
	ins := ioc.instanceMap[name]
	if ins != nil {
		return errors.New("instance [" + name + "] is exists")
	}
	ioc.instanceMap[name] = x
	fmt.Println("ioc load --> ", name)
	return nil
}

func (ioc *SpringIOC) GetByName(name string) interface{} {
	return ioc.instanceMap[name]
}

func (ioc *SpringIOC) instanceName(p reflect.Type) string {
	return p.Elem().PkgPath() + "/" + p.Elem().Name()
}

func (ioc *SpringIOC) createIns(x interface{}) interface{} {
	refType := reflect.TypeOf(x)
	refVal := reflect.ValueOf(x)
	for i := 0; i < refType.Elem().NumField(); i++ {
		field := refType.Elem().Field(i)
		value := refVal.Elem().FieldByName(field.Name)
		valTag := field.Tag.Get("val")
		if valTag != "" {
			ioc.injectValue(valTag, value)
			continue
		}
		iocTag := field.Tag.Get("ioc")
		if iocTag == "" || iocTag != "auto" || field.Type.Kind() != reflect.Ptr {
			continue
		}
		ins := ioc.GetByName(ioc.instanceName(field.Type))
		if ins != nil {
			value.Set(reflect.ValueOf(ins))
		} else {
			newIns := reflect.New(field.Type.Elem())
			value.Set(newIns)
			newInterface := newIns.Interface()
			ioc.setIns(ioc.instanceName(field.Type), newInterface)
			ioc.createIns(newInterface)
		}
	}
	return x
}

func (ioc *SpringIOC) injectValue(key string, value reflect.Value) error {
	compile := regexp.MustCompile("^\\$\\{(.+?)\\}$")
	if !compile.MatchString(key) {
		return nil
	}
	chain := compile.FindAllStringSubmatch(key, -1)[0][1]
	data, err := ioc.config.ChainCall(chain)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("inject value fail, " + key + " is not found")
	}
	switch value.Kind() {
	case reflect.String:
		value.SetString(data.(string))
	case reflect.Int:
		value.SetInt(int64(data.(float64)))
	case reflect.Bool:
		value.SetBool(data.(bool))
	}
	return nil
}
