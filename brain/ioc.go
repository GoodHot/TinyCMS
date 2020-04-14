package brain

import (
	"errors"
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/wang22/easyjson"
	"reflect"
	"regexp"
	"strconv"
)

type BrainIOC struct {
	instanceMap map[string]interface{}
	config      *easyjson.EasyJSON
	log         Logger
}

func (ioc *BrainIOC) Reg(x interface{}) (ins interface{}, err error) {
	insType := reflect.TypeOf(x)
	if insType.Kind() != reflect.Ptr {
		ioc.log.Error("Register must be struct")
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

func (ioc *BrainIOC) RegInterface(i interface{}, x interface{}) error {
	interfaceType := reflect.TypeOf(i)
	if interfaceType.Elem().Kind() != reflect.Interface {
		msg := strs.Fmt("%v is not interface", interfaceType)
		ioc.log.Error(msg)
		return errors.New(msg)
	}
	interfaceName := ioc.instanceName(interfaceType)
	fmt.Println("name", interfaceName)
	ins := ioc.GetByName(interfaceName)
	if ins != nil {
		return nil
	}
	ioc.setIns(interfaceName, x)
	ioc.createIns(x)
	return nil
}

func (ioc *BrainIOC) Get(x interface{}) interface{} {
	refType := reflect.TypeOf(x)
	return ioc.GetByName(ioc.instanceName(refType))
}

func (ioc *BrainIOC) setIns(name string, x interface{}) error {
	ins := ioc.GetByName(name)
	if ins != nil {
		ioc.log.Error("instance [%s] is exists", name)
		return errors.New("instance [" + name + "] is exists")
	}
	ioc.instanceMap[name] = x
	ioc.log.Info("IOC loading --> %s", name)
	return nil
}

func (ioc *BrainIOC) GetByName(name string) interface{} {
	return ioc.instanceMap[name]
}

func (ioc *BrainIOC) instanceName(p reflect.Type) string {
	return p.Elem().PkgPath() + "/" + p.Elem().Name()
}

func (ioc *BrainIOC) createIns(x interface{}) interface{} {
	refType := reflect.TypeOf(x)
	refVal := reflect.ValueOf(x)
	for i := 0; i < refType.Elem().NumField(); i++ {
		field := refType.Elem().Field(i)
		value := refVal.Elem().FieldByName(field.Name)
		valTag := field.Tag.Get("val")
		if valTag != "" {
			err := ioc.injectValue(valTag, value)
			if err != nil {
				panic(errors.New(refType.Elem().PkgPath() + "/" + refType.Elem().Name() + "." + field.Name + " [" + err.Error() + "]"))
			}
			continue
		}
		iocTag := field.Tag.Get("ioc")
		if iocTag == "" || iocTag != "auto" || (field.Type.Kind() != reflect.Ptr && field.Type.Kind() != reflect.Interface) {
			continue
		}
		var insName string
		if field.Type.Kind() == reflect.Ptr {
			insName = ioc.instanceName(field.Type)
		} else {
			insName = field.Type.PkgPath() + "/" + field.Type.Name()
		}
		ins := ioc.GetByName(insName)
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
	method := refVal.MethodByName("Startup")
	if method.Kind() == reflect.Func {
		err := method.Call(nil)[0]
		if !err.IsNil() {
			msg := err.Interface().(error).Error()
			errMsg := refType.Elem().PkgPath() + "/" + refType.Elem().Name() + " -> call Startup function error, error msg:" + msg
			ioc.log.Error(errMsg)
			panic(errors.New(errMsg))
		}
	}
	return x
}

func (ioc *BrainIOC) injectValue(key string, value reflect.Value) error {
	compile := regexp.MustCompile("^\\$\\{(.+?)\\}$")
	if !compile.MatchString(key) {
		switch value.Kind() {
		case reflect.String:
			value.SetString(key)
		case reflect.Int:
			if v, e := strconv.Atoi(key); e != nil {
				return e
			} else {
				value.SetInt(int64(v))
			}
		case reflect.Bool:
			value.SetBool(key == "true")
		}
		return nil
	}
	chain := compile.FindAllStringSubmatch(key, -1)[0][1]
	data, err := ioc.config.ChainCall(chain)
	if err != nil {
		return err
	}
	if data == nil {
		ioc.log.Error("inject value fail, %s is not found", key)
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
