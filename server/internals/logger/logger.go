package utils

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

/*
InitializeZapCustomLogger Funtion initializes a logger using uber-go/zap package in the application.
*/
func InitializeZapCustomLogger() {
	//conf := zap.Config{
	//	Encoding:    "json",
	//	Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
	//	OutputPaths: []string{viper.GetString("logger-output-path"), "stdout"},
	//	EncoderConfig: zapcore.EncoderConfig{
	//		LevelKey:     "level",
	//		TimeKey:      "time",
	//		CallerKey:    "file",
	//		MessageKey:   "msg",
	//		EncodeLevel:  zapcore.LowercaseLevelEncoder,
	//		EncodeTime:   zapcore.ISO8601TimeEncoder,
	//		EncodeCaller: zapcore.ShortCallerEncoder,
	//	},
	//}
	//
	//Log, _ = conf.Build()
	Log, _ = zap.NewProduction()
}