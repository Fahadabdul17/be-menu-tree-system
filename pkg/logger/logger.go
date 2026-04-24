package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = "" // Disable stacktrace for cleaner logs

	var err error
	Log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func Get() *zap.Logger {
	if Log == nil {
		Init()
	}
	return Log
}
