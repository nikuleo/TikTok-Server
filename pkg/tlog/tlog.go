package tlog

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	once   sync.Once
)

var (
	Debug func(msg string, fields ...zap.Field)
	Info  func(msg string, fields ...zap.Field)
	Warn  func(msg string, fields ...zap.Field)
	Error func(msg string, fields ...zap.Field)
	Fatal func(msg string, fields ...zap.Field)

	// wrap sugar
	Debugf func(template string, args ...interface{})
	Infof  func(template string, args ...interface{})
	Warnf  func(template string, args ...interface{})
	Errorf func(template string, args ...interface{})
	Fatalf func(template string, args ...interface{})

	String   = zap.String
	Int      = zap.Int
	Err      = zap.Error
	Any      = zap.Any
	Duration = zap.Duration
)

func InitLog() {
	once.Do(func() {
		logger = logConfig()
		sugar = logger.Sugar()
	})

	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal

	Debugf = sugar.Debugf
	Infof = sugar.Infof
	Warnf = sugar.Warnf
	Errorf = sugar.Errorf
	Fatalf = sugar.Fatalf

}

func logConfig() *zap.Logger {

	homePath := os.Getenv("HOME")
	// 日志文件分割配置
	fileWriterHook := &lumberjack.Logger{
		Filename:   homePath + "/var/log/tiktokserver/server.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	// 日志文件输出配置
	fileEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,                        // 全大写日志等级标识
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 终端输出配置
	stdEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	stdEncoder := zapcore.NewConsoleEncoder(stdEncoderConfig)

	fileWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(fileWriterHook))
	stdWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))

	// 日志级别过滤
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return lvl >= zapcore.InfoLevel
	// })

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, warnLevel),
		zapcore.NewCore(stdEncoder, stdWriter, debugLevel),
	)

	if os.Getenv("GIN_DEBUG") == "true" {
		caller := zap.AddCaller()
		development := zap.Development()
		return zap.New(core, caller, development)
	} else {
		return zap.New(core)
	}
}

func Sync() {
	logger.Sync()
}
