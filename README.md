# logx

封装了Zap和lumberjack的日志库


```go
func TestLoger(t *testing.T) {
	logOpts := &logx.Options{
		//log path 日志文件路径
		LogFile: "logs/test.log",
		//log size 日志文件大小，单位：MB
		MaxSize: 10,
		//log age 日志文件保存时间，单位：天
		MaxAge: 30,
		//log backups 日志文件备份数量
		MaxBackups: 10,
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
{"level":"info","ts":"2025-05-04T11:26:13.666+0800","caller":"test/loger_test.go:27","msg":"logx: error"}

{"level":"info","ts":"2025-05-04T11:26:13.666+0800","caller":"test/loger_test.go:30","msg":"logx","info":"test: error"}
```