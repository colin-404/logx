# logx

封装了Zap和lumberjack的日志库


```go
package test

import (
	"errors"
	"testing"

	"github.com/colin-404/logx"
)

func TestLoger(t *testing.T) {
	logOpts := &logx.Options{
		//log path 日志文件路径,默认：./default.log
		LogFile: "logs/test.log",
		//log size 日志文件大小，单位：MB,默认：5
		MaxSize: 10,
		//log age 日志文件保存时间，单位：天,默认：3
		MaxAge: 30,
		//log backups 日志文件备份数量,默认：3
		MaxBackups: 10,

		//time format 日志时间格式,默认：EpochNanos
		TimeFormat: logx.TimeFormats.EpochNanos,
	}
	loger := logx.NewLoger(logOpts)
	logx.InitLogger(loger)

	err := errors.New("error")

	// info
	logx.Infof("logx: %v", err)

	// add msg to info log
	logx.Infomf("logx", "test: %v", err)
}


```

日志格式
```
{"level":"info","ts":1746334560832477499,"caller":"test/loger_test.go:28","msg":"logx: error"}
{"level":"info","ts":1746334560832745842,"caller":"test/loger_test.go:31","msg":"logx","info":"test: error"}
```