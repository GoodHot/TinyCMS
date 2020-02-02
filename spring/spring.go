package spring

type Spring struct {
	instanceMap map[string]interface{}
}

func (s *Spring) Reg(i interface{}) interface{} {
	return i
}

func (*Spring) Get(i interface{}) {

}

var ctx *Spring

func AppCtx() *Spring {
	return ctx
}

func init() {
	ctx = &Spring{}
}
