package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var (
	zapSugared *zap.SugaredLogger

	levelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
)

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func getSyncWritter(fileName string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
}

func getFileLogger() *zap.SugaredLogger {
	fileName := "log.log"
	level := getLoggerLevel("debug")
	syncWriter := getSyncWritter(fileName)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}

func getProductionLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}

func InitLogger(isFile bool) {
	if isFile {
		zapSugared = getFileLogger()
	} else {
		zapSugared = getProductionLogger()
	}
}

func Debug(args ...interface{}) {
	zapSugared.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapSugared.Debugf(template, args...)
}

func Info(args ...interface{}) {
	zapSugared.Info(args...)
}
func Log(args ...interface{}) {
	zapSugared.Info(args...)
}
func Logf(template string, args ...interface{}) {
	zapSugared.Infof(template, args...)
}
func Infof(template string, args ...interface{}) {
	zapSugared.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zapSugared.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapSugared.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zapSugared.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapSugared.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	zapSugared.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	zapSugared.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	zapSugared.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	zapSugared.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	zapSugared.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	zapSugared.Fatalf(template, args...)
}
