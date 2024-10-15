package logger

import (
	"backend-go/pkg/setting"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Level

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger {
		Filename:   config.Filename,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, //days
		Compress:   config.Compress, // disabled by default
	}

	core := zapcore.NewCore(
		encoder, 
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), 
		level)
	// logger := zap.New(core, zap.AddCaller())

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format logs message
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1728893847.5158699 -> 2024-10-14T15:17:27.515+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// timestamp -> time
	encoderConfig.TimeKey = "time"

	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:23"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
