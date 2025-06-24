package log

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugarLogger *zap.SugaredLogger
)

func init() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	loggerConfig.Development = true
	loggerConfig.Level.SetLevel(zap.DebugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	sugarLogger = logger.Sugar()
	sugarLogger.Info("Logger initialized")
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}
func Debugln(args ...interface{}) {
	sugarLogger.Debugln(args...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Debugw(msg, keysAndValues...)
}
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}
func Errorln(args ...interface{}) {
	sugarLogger.Errorln(args...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Errorw(msg, keysAndValues...)
}
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}
func Infoln(args ...interface{}) {
	sugarLogger.Infoln(args...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	sugarLogger.Infow(msg, keysAndValues...)
}
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}
func Warnln(args ...interface{}) {
	sugarLogger.Warnln(args...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Warnw(msg, keysAndValues...)
}
