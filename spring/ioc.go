package spring

import (
	"errors"
	"fmt"
	"reflect"
)

type SpringIOC struct {
	instanceMap map[string]interface{}
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
			switch value.Kind() {
			case reflect.String:
				value.SetString("fffffffffff")
			}
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
