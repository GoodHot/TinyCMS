package logger

import (
	"flag"
	"fmt"
)

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

type TinyLogger struct {
	Level  string `val:"${logger.level}"`
	Output string `val:"${logger.output}"`
}

func (t *TinyLogger) Startup() error {
	//lf, err := os.OpenFile(t.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	//if err != nil {
	//	logger.Fatalf("Failed to open log file: %v", err)
	//}
	//defer lf.Close()
	//t.logger = logger.Init("LoggerExample", *verbose, true, lf)
	return nil
}

func (t *TinyLogger) Info(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
func (t *TinyLogger) Warn(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
func (t *TinyLogger) Error(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
func (t *TinyLogger) Debug(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
func (t *TinyLogger) Trace(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
func (t *TinyLogger) Fatal(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
func (t *TinyLogger) Panic(msg string, param ...interface{}) {
	fmt.Printf(msg + "\n", param...)
}
