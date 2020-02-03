package spring

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

func init() {
	ctx = &Spring{
		ioc: &SpringIOC{
			instanceMap: make(map[string]interface{}),
		},
	}
}
