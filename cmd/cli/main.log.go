package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "tung-lee", 21) // similar to fmt.Printf

	// logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "tung-lee"), zap.Int("age", 21))

	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	encoder := getEncoderLog()
	sync:= getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info Log", zap.Int("info-log", 1))
	logger.Error("Error Log", zap.Int("error-log", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	// file, _ := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	file, _ := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
