package test

import (
	"github.com/GoodHot/TinyCMS/controller"
	"reflect"
	"testing"
)

func TestNewStruct(t *testing.T) {
	i := &controller.Controller{}
	to := reflect.TypeOf(i)
	vo := reflect.ValueOf(i)
	for i := 0; i < to.Elem().NumField(); i++ {
		field := to.Elem().Field(i)
		val := vo.Elem().FieldByName(field.Name)
		if field.Type.Kind() != reflect.Ptr {
			continue
		}
		iocName, ok := field.Tag.Lookup("ioc")
		if !ok || iocName != "auto" {
			continue
		}
		ins := reflect.New(field.Type.Elem())
		ff := ins.Interface()
		tt := reflect.TypeOf(ff)
		vv := reflect.ValueOf(ff)
		for j := 0; j < tt.Elem().NumField(); j++ {
			fv := vv.Elem().Field(j)
			if fv.Kind() == reflect.String {
				fv.SetString("ooooooooooo")
			}
			//ft := tt.Elem().Field(j)
			//if ft.Type.Kind() == reflect.String {
			//	fv := vv.Elem().FieldByName(j)
			//	fmt.Println(fv)
			//}
			//fv := vv.Elem().FieldByName(ft.Name)
			//fmt.Println(fv.Elem())
		}
		//fmt.Printf("%v", ins)
		//vv := reflect.TypeOf(ins.Interface())
		////fmt.Println(vv.Elem().Kind())
		val.Set(ins)
	}
	//fmt.Println(i.Config.Env)
}
