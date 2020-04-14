package logger

import (
	"github.com/labstack/gommon/log"
)

type TinyLogger struct {
	Level string `val:"${logger.level}"`
}

func (t *TinyLogger) Info(fmt string, param ...interface{}) {
	log.Infof(fmt, param...)
}
func (*TinyLogger) Warn(fmt string, param ...interface{}) {
	log.Warnf(fmt, param...)
}
func (*TinyLogger) Error(fmt string, param ...interface{}) {
	log.Errorf(fmt, param...)
}
func (*TinyLogger) Debug(fmt string, param ...interface{}) {
	log.Debugf(fmt, param...)
}
func (*TinyLogger) Trace(fmt string, param ...interface{}) {
	log.Debugf(fmt, param...)
}
func (*TinyLogger) Fatal(fmt string, param ...interface{}) {
	log.Fatalf(fmt, param...)
}
func (*TinyLogger) Panic(fmt string, param ...interface{}) {
	log.Panicf(fmt, param...)
}
