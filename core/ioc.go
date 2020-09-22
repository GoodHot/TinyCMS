package core

import (
	"errors"
	"fmt"
	"reflect"
)

type IOC struct {
	cfg    *Config
	main   interface{}
	insMap map[string]interface{}
}

func (ioc *IOC) Register(ins interface{}) {
	if ioc.insMap == nil {
		ioc.insMap = make(map[string]interface{})
	}
	insName := ioc.instanceName(ins)
	if ioc.Get(insName) != nil {
		return
	}
	ioc.createIns(ins)
}
func (ioc *IOC) createIns(ins interface{}) {
	refVal := reflect.ValueOf(ins)
	refType := reflect.TypeOf(ins)
	fmt.Println("ddd", refVal.Elem().NumField())
	for i := 0; i < refVal.Elem().NumField(); i++ {
		field := refType.Elem().Field(i)
		iocTag := field.Tag.Get("ioc")
		valTag := field.Tag.Get("val")
		if valTag != "" {
			fmt.Println(valTag)
		} else if iocTag != "" {

		} else {
			continue
		}
	}

	//refType := reflect.TypeOf(x)
	//refVal := reflect.ValueOf(x)
	//for i := 0; i < refType.Elem().NumField(); i++ {
	//	field := refType.Elem().Field(i)
	//	value := refVal.Elem().FieldByName(field.Name)
	//	valTag := field.Tag.Get("val")
	//	if valTag != "" {
	//		err := ioc.injectValue(valTag, value)
	//		if err != nil {
	//			panic(errors.New(refType.Elem().PkgPath() + "/" + refType.Elem().Name() + "." + field.Name + " [" + err.Error() + "]"))
	//		}
	//		continue
	//	}
	//	iocTag := field.Tag.Get("ioc")
	//	if iocTag == "" || iocTag != "auto" || (field.Type.Kind() != reflect.Ptr && field.Type.Kind() != reflect.Interface) {
	//		continue
	//	}
	//	var insName string
	//	if field.Type.Kind() == reflect.Ptr {
	//		insName = ioc.instanceName(field.Type)
	//	} else {
	//		insName = field.Type.PkgPath() + "/" + field.Type.Name()
	//	}
	//	ins := ioc.Get(insName)
	//	if ins != nil {
	//		value.Set(reflect.ValueOf(ins))
	//	} else {
	//		newIns := reflect.New(field.Type.Elem())
	//		value.Set(newIns)
	//		newInterface := newIns.Interface()
	//		ioc.setIns(ioc.instanceName(field.Type), newInterface, false)
	//		ioc.createIns(newInterface)
	//	}
	//}
	//method := refVal.MethodByName("Startup")
	//if method.Kind() == reflect.Func {
	//	err := method.Call(nil)[0]
	//	if !err.IsNil() {
	//		msg := err.Interface().(error).Error()
	//		errMsg := refType.Elem().PkgPath() + "/" + refType.Elem().Name() + " -> call Startup function error, error msg:" + msg
	//		panic(errors.New(errMsg))
	//	}
	//}
}

func (ioc *IOC) injectValue(key string, value reflect.Value) error {
	//compile := regexp.MustCompile("^\\$\\{(.+?)\\}$")
	//if !compile.MatchString(key) {
	//	switch value.Kind() {
	//	case reflect.String:
	//		value.SetString(key)
	//	case reflect.Int:
	//		if v, e := strconv.Atoi(key); e != nil {
	//			return e
	//		} else {
	//			value.SetInt(int64(v))
	//		}
	//	case reflect.Bool:
	//		value.SetBool(key == "true")
	//	}
	//	return nil
	//}
	//chain := compile.FindAllStringSubmatch(key, -1)[0][1]
	//data, err := ioc.config.ChainCall(chain)
	//if err != nil {
	//	return err
	//}
	//if data == nil {
	//	ioc.log.Error("inject value fail, %s is not found", key)
	//	return errors.New("inject value fail, " + key + " is not found")
	//}
	//switch value.Kind() {
	//case reflect.String:
	//	value.SetString(data.(string))
	//case reflect.Int:
	//	value.SetInt(int64(data.(float64)))
	//case reflect.Bool:
	//	value.SetBool(data.(bool))
	//}
	return nil
}

func (ioc *IOC) instanceName(ins interface{}) string {
	insType := reflect.TypeOf(ins)
	return insType.Elem().PkgPath() + "/" + insType.Elem().Name()
}

func (ioc *IOC) Init(main interface{}, cfgPath string) error {
	if main == nil {
		return errors.New("main is nil")
	}
	ioc.main = main
	ioc.cfg = new(Config)
	err := ioc.cfg.ReadConfig(cfgPath)
	if err != nil {
		return err
	}
	ioc.Register(main)
	return nil
}

func (ioc *IOC) GetMain() interface{} {
	return ioc.main
}

func (ioc *IOC) Get(name string) interface{} {
	return ioc.insMap[name]
}
