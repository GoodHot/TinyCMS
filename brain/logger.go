package brain

type Logger interface {
	Info(log string, param ...interface{})
	Warn(log string, param ...interface{})
	Error(log string, param ...interface{})
	Debug(log string, param ...interface{})
	Trace(log string, param ...interface{})
	Fatal(log string, param ...interface{})
	Panic(log string, param ...interface{})
}
