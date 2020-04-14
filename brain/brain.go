package brain

import (
	"github.com/wang22/easyjson"
	"io/ioutil"
	"reflect"
)

type Brain struct {
	ioc *BrainIOC
}

func (b *Brain) Register(i interface{}) interface{} {
	ins, _ := b.ioc.Reg(i)
	return ins
}

func (b *Brain) SetLogger(logger Logger) {
	logger.Info("Brain set logger %v", reflect.TypeOf(logger))
	b.ioc.log = logger
	b.ioc.RegInterface((*Logger)(nil), logger)
}

func Load(cfgFile string) (*Brain, error) {
	fileData, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}
	ej := easyjson.ParseJSON(fileData)
	return &Brain{
		ioc: &BrainIOC{
			instanceMap: make(map[string]interface{}),
			config:      ej,
		},
	}, nil
}
