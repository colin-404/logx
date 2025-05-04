# logx

封装了Zap和lumberjack的日志库


```
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

logx.Infof("logx", "test")

```

日志格式
```
{"level":"info","ts":"2025-05-04T11:02:15.248+0800","caller":"test/loger_test.go:23","msg":"logx","info":"test"}
```