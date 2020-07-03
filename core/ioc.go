package core

type IoC interface {
	Reg(i interface{})
	Singleton()
}
