package core

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
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
	for i := 0; i < refVal.Elem().NumField(); i++ {
		field := refType.Elem().Field(i)
		iocTag := field.Tag.Get("ioc")
		valTag := field.Tag.Get("val")
		if valTag != "" {
			if err := ioc.injectValue(valTag, refVal.Elem().FieldByName(field.Name)); err != nil {
				panic(errors.New(refType.Elem().PkgPath() + "/" + refType.Elem().Name() + "." + field.Name + " [" + err.Error() + "]"))
			}
		} else if iocTag != "" {

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

var injectValueRegex = regexp.MustCompile("^\\$\\{(.+?)\\}$")

func (ioc *IOC) injectValue(key string, value reflect.Value) error {
	var tmpVal interface{}
	if !injectValueRegex.MatchString(key) {
		tmpVal = key
	} else {
		match := injectValueRegex.FindAllStringSubmatch(key, -1)
		val, exists := ioc.cfg.GetVal(match[0][1])
		if !exists {
			return errors.New("inject value fail, " + key + " is not found")
		}
		tmpVal = val
	}
	valKind := reflect.TypeOf(tmpVal).Kind()
	switch value.Kind() {
	case reflect.String:
		value.SetString(tmpVal.(string))
	case reflect.Int:
		switch valKind {
		case reflect.Int:
			value.SetInt(int64(tmpVal.(int)))
		case reflect.Int64:
			value.SetInt(tmpVal.(int64))
		case reflect.Float32:
			value.SetInt(int64(tmpVal.(float32)))
		case reflect.Float64:
			value.SetInt(int64(tmpVal.(float64)))
		case reflect.String:
			if v, err := strconv.Atoi(tmpVal.(string)); err != nil {
				panic("inject value fail, can not convert " + tmpVal.(string) + " to int")
			} else {
				value.SetInt(int64(v))
			}
		default:
			panic("inject value fail, can not convert " + valKind.String() + " to bool")
		}
	case reflect.Bool:
		if valKind == reflect.Bool {
			value.SetBool(tmpVal.(bool))
		} else if valKind == reflect.String {
			value.SetBool(tmpVal.(string) == "true")
		} else {
			panic("inject value fail, can not convert " + valKind.String() + " to bool")
		}
	}
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