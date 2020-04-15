package brain

import "fmt"

type Logger interface {
	Info(log string, param ...interface{})
	Warn(log string, param ...interface{})
	Error(log string, param ...interface{})
	Debug(log string, param ...interface{})
	Trace(log string, param ...interface{})
	Fatal(log string, param ...interface{})
	Panic(log string, param ...interface{})
}

type StdLogger struct {
}

func (*StdLogger) Info(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
func (*StdLogger) Warn(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
func (*StdLogger) Error(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
func (*StdLogger) Debug(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
func (*StdLogger) Trace(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
func (*StdLogger) Fatal(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
func (*StdLogger) Panic(log string, param ...interface{}) {
	fmt.Printf(log, param...)
}
