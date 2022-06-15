// 
// @Author: jason zhou
// @Description: 
// @File:  logger.go
// @Version: 0.0.1
// @Date: 2022/6/7 16:32
// 

package log

import (
    "fmt"
)

//
//
//
type Logger interface {
    Info(v ...interface{})
    Error(v ...interface{})
}

var logger Logger = NewStdLogger()

//
//
//
func InitLogger(log Logger) {
    if logger == nil {
        panic("logger is nil")
    }
    logger = log
}

//
//
//
func GetLogger() Logger {
    return logger
}

//
//
//
type StdLogger struct {
}

//
//
//
func NewStdLogger() *StdLogger {
    return &StdLogger{}
}

//
//
//
func (_this *StdLogger) Info(v ...interface{}) {
    fmt.Println(v)
}

//
//
//
func (_this *StdLogger) Error(v ...interface{}) {
    fmt.Println(v)
}
