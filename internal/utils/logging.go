package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func init() {

	//Caller  pc, := runtime.FuncforPC(3)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "message",
		StacktraceKey:  "stackTree",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logfile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err.Error())
	}
	fileSyncer := zapcore.AddSync(logfile)
	consoleSycner := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileSyncer, zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleSycner, zapcore.InfoLevel),
	)
	globalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	globalLogger.Info("Logger init succeed")
}

func GetLogger() *zap.Logger {
	return globalLogger
}
