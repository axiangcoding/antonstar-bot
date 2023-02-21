package logging

import (
	"github.com/axiangcoding/antonstar-bot/setting"
	"log"
	"os"
	"path"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	_logger      = zap.NewNop()
	_sugarLogger = _logger.Sugar()
)

func getLogLevel(level string) zapcore.Level {
	level = strings.ToUpper(level)
	var zapLevel zapcore.Level
	switch level {
	case "DEBUG":
		zapLevel = zapcore.DebugLevel
	case "INFO":
		zapLevel = zapcore.InfoLevel
	case "WARN":
		zapLevel = zapcore.WarnLevel
	case "ERROR":
		zapLevel = zapcore.ErrorLevel
	case "FATAL":
		zapLevel = zapcore.FatalLevel
	default:
		log.Fatalln("no such log level")
	}
	return zapLevel
}

func getEncoder(enc string) zapcore.Encoder {
	var encoder zapcore.Encoder
	encConf := zap.NewProductionEncoderConfig()
	// 可读性时间戳
	encConf.EncodeTime = zapcore.ISO8601TimeEncoder
	// 将日志等级大写
	encConf.EncodeLevel = zapcore.CapitalLevelEncoder
	// 调用者命名
	encConf.EncodeCaller = zapcore.FullCallerEncoder
	switch enc {
	case setting.AppLogFileEncoderJson:
		encoder = zapcore.NewJSONEncoder(encConf)
	case setting.AppLogFileEncoderConsole:
		encConf.EncodeCaller = zapcore.FullCallerEncoder
		encConf.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encConf)
	default:
		log.Fatalln("no such log file encoder")
	}
	return encoder
}

func InitLogger(level string, mode string) {
	logLevel := getLogLevel(level)

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: path.Join(setting.C().App.Log.File.Dir,
			"application.log"),
		MaxSize:    100, // megabytes
		MaxBackups: 100,
		MaxAge:     60, // days
		Compress:   true,
	})
	encoder := getEncoder(setting.C().App.Log.File.Encoder)

	var cores []zapcore.Core
	prodCore := zapcore.NewCore(
		encoder,
		w,
		logLevel,
	)
	cores = append(cores, prodCore)
	if mode == setting.AppRunModeDebug {
		debugCore := zapcore.NewCore(getEncoder("console"), os.Stdout, logLevel)
		cores = append(cores, debugCore)
	}

	tee := zapcore.NewTee(cores...)
	_logger = zap.New(tee).WithOptions(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(_logger)
}

func S() *zap.SugaredLogger {
	return _sugarLogger
}

func L() *zap.Logger {
	return _logger
}

func Any(key string, value any) zapcore.Field {
	return zap.Any(key, value)
}

func Error(err error) zapcore.Field {
	return zap.Error(err)
}

func Errors(key string, errors []error) zapcore.Field {
	return zap.Errors(key, errors)
}
