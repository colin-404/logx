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
	defaultTimeFormat = "EpochNanos"
	defaultCaller     = true
)

// TimeFormatsStruct provides constants as struct fields for easier access.
type timeFormatContainer struct {
	ISO8601     string
	RFC3339     string
	EpochMillis string
	EpochNanos  string
	Epoch       string
}

// TimeFormats is an instance of the container holding the format constants.
var TimeFormats = timeFormatContainer{
	ISO8601:     "ISO8601",
	RFC3339:     "RFC3339",
	EpochMillis: "EpochMillis",
	EpochNanos:  "EpochNanos",
	Epoch:       "Epoch",
}

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

	if opts.TimeFormat == "" {
		opts.TimeFormat = defaultTimeFormat
	}

	// Handle default for Caller pointer
	if opts.Caller == nil {
		callerDefault := defaultCaller // Create a local variable
		opts.Caller = &callerDefault   // Take address of the variable
	}

	// Determine the effective caller setting
	enableCaller := defaultCaller
	if opts.Caller != nil {
		enableCaller = *opts.Caller
	}

	hook := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  false,
		Compress:   false,
	}

	encoderConfig := zap.NewProductionEncoderConfig()

	// caller 是否显示调用者
	if enableCaller { // Use the effective setting
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoderConfig.CallerKey = "source"
	}
	encoderConfig.TimeKey = "timestamp"

	// Determine the time encoder based on opts.TimeFormat
	switch opts.TimeFormat {
	case "ISO8601":
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	case "RFC3339":
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	case "EpochMillis":
		encoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	case "EpochNanos":
		encoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
	case "Epoch":
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(hook)),
		zapcore.Level(logLevel))

	// Conditionally add caller options to zap.New
	var zapOptions []zap.Option
	if enableCaller {
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1))
	}

	logger := zap.New(core, zapOptions...)
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

func Fatalf(msg string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(msg+"\n", v...)
		return
	}
	if defaultLogger.lvl <= FatalLevel {
		// defaultLogger.provider.Fatal(msg, zap.String("info", fmt.Sprintf(format, v...)))
		defaultLogger.provider.Fatal(fmt.Sprintf(msg, v...))
	}
}

func Fatalmf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= FatalLevel {
		defaultLogger.provider.Fatal(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Errorf(msg string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(msg+"\n", v...)
		return
	}
	if defaultLogger.lvl <= ErrorLevel {
		defaultLogger.provider.Error(fmt.Sprintf(msg, v...))
	}
}

func Errormf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= ErrorLevel {
		defaultLogger.provider.Error(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Warnf(msg string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(msg+"\n", v...)
		return
	}
	if defaultLogger.lvl <= WarnLevel {
		defaultLogger.provider.Warn(fmt.Sprintf(msg, v...))
	}
}

func Warnmf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= WarnLevel {
		defaultLogger.provider.Warn(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Info(v interface{}) {
	if defaultLogger == nil {
		fmt.Printf("%v"+"\n", v)
		return
	}
	if defaultLogger.lvl <= InfoLevel {
		defaultLogger.provider.Info(fmt.Sprintf("%v", v))
	}
}

func Infof(msg string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(msg+"\n", v...)
		return
	}
	if defaultLogger.lvl <= InfoLevel {
		defaultLogger.provider.Info(fmt.Sprintf(msg, v...))
	}
}

// add msg to log
func Infomf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= InfoLevel {
		defaultLogger.provider.Info(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}

func Debug(v interface{}) {
	if defaultLogger == nil {
		fmt.Printf("%v"+"\n", v)
		return
	}
	if defaultLogger.lvl <= DebugLevel {
		defaultLogger.provider.Debug(fmt.Sprintf("%v", v))
	}
}

func Debugf(msg string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(msg+"\n", v...)
		return
	}
	if defaultLogger.lvl <= DebugLevel {
		defaultLogger.provider.Debug(fmt.Sprintf(msg, v...))
	}
}

func Debugmf(msg string, format string, v ...interface{}) {
	if defaultLogger == nil {
		fmt.Printf(format+"\n", v...)
		return
	}
	if defaultLogger.lvl <= DebugLevel {
		defaultLogger.provider.Debug(msg, zap.String("info", fmt.Sprintf(format, v...)))
	}
}
