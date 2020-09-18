package core

import "errors"

type IOC struct {
	cfg *Config
	main interface{}
}

func (*IOC) Register(i interface{}) {
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

	return nil
}

func (ioc *IOC) GetMain() interface{} {
	return ioc.main
}

func (*IOC) Get(name string) {
}
