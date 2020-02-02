package spring

type Spring struct {
}

func (s *Spring) Reg(i interface{}) interface{} {
	return i
}

func (*Spring) Get(name string) {

}

var ctx *Spring

func AppCtx() *Spring {
	return ctx
}

func init() {
	ctx = &Spring{}
}
