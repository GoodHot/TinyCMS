package brain

import (
	"github.com/wang22/easyjson"
	"io/ioutil"
)

type Brain struct {
	ioc *BrainIOC
}

func (s *Brain) Reg(i interface{}) interface{} {
	ins, _ := s.ioc.Reg(i)
	return ins
}

func (*Brain) Get(i interface{}) {

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
