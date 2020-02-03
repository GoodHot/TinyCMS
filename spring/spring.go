package spring

import (
	"github.com/wang22/easyjson"
	"io/ioutil"
)

type Spring struct {
	ioc *SpringIOC
}

func (s *Spring) Reg(i interface{}) interface{} {
	ins, _ := s.ioc.Reg(i)
	return ins
}

func (*Spring) Get(i interface{}) {

}

var ctx *Spring

func AppCtx() *Spring {
	return ctx
}

func Load(cfgFile string) error {
	fileData, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return err
	}
	ej := easyjson.ParseJSON(fileData)
	ctx = &Spring{
		ioc: &SpringIOC{
			instanceMap: make(map[string]interface{}),
			config:      ej,
		},
	}
	return nil
}
