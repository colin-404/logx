package test

import (
	"errors"
	"testing"

	"github.com/colin-404/logx"
)

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
