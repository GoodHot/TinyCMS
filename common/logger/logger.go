package logger

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
)

type TinyLogger struct {
	Level string `val:"${logger.level}"`
}

func (*TinyLogger) Info(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
func (*TinyLogger) Warn(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
func (*TinyLogger) Error(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
func (*TinyLogger) Debug(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
func (*TinyLogger) Trace(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
func (*TinyLogger) Fatal(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
func (*TinyLogger) Panic(log string, param ...interface{}) {
	fmt.Println(strs.Fmt(log, param...))
}
