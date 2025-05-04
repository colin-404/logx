package logx

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultLogFile    = "./default.log"
	defaultMaxSize    = 5
	defaultMaxAge     = 3
	defaultMaxBackups = 3
)

var defaultLogger *Loger

func InitLogger(logger *Loger) {
	defaultLogger = logger
}

// Logger
type Loger struct {
	provider *zap.Logger
	msg      string
	lvl      int
}

func NewLoger(opts *Options) *Loger {

	logFile := defaultLogFile
	if opts.LogFile != "" {
		logFile = opts.LogFile
	}
	maxSize := defaultMaxSize
	if opts.MaxSize != 0 {
		maxSize = opts.MaxSize
	}
	maxAge := defaultMaxAge
	if opts.MaxAge != 0 {
		maxAge = opts.MaxAge
	}
	maxBackups := defaultMaxBackups
	if opts.MaxSize != 0 {
		maxBackups = opts.MaxBackups
	}
	logLevel := opts.Level

	hook := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  false,
		Compress:   false,
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(hook)),
		zapcore.Level(logLevel))
	// return &Loger{provider: zap.New(core), lvl: logLevel}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &Loger{provider: logger, lvl: logLevel}
}

func (l *Loger) SetMsg(msg string) {
	l.msg = msg
}

func (l *Loger) Print(v ...interface{}) {
	l.provider.Info(l.msg, zap.Any("info", v))
}

func (l *Loger) Printf(format string, v ...interface{}) {
	l.provider.Info(l.msg, zap.Any("info", fmt.Sprintf(format, v...)))
}

func (l *Loger) Println(v ...interface{}) {
	l.provider.Info(l.msg, zap.Any("info", v))
}

func Fatalf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= FatalLevel {
		defaultLogger.provider.Fatal(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Errorf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= ErrorLevel {
		defaultLogger.provider.Error(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Warnf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= WarnLevel {
		defaultLogger.provider.Warn(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Infof(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= InfoLevel {
		defaultLogger.provider.Info(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Debugf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= DebugLevel {
		defaultLogger.provider.Debug(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}
