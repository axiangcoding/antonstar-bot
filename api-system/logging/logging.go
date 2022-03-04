package logging

import (
	"axiangcoding/antonstar/api-system/settings"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logFile *zap.SugaredLogger
var logConsole *zap.SugaredLogger

var enableFileLog = false

func Setup() {
	enableFileLog = settings.Config.App.Log.File.Enable
	logger, _ := zap.NewDevelopment()
	logConsole = logger.Sugar()
	// 是否打印日志到文件中
	if enableFileLog {
		// 设置application的日志输出
		zapLevel := zapcore.InfoLevel
		level := settings.Config.App.Log.Level
		switch level {
		case "info":
			zapLevel = zapcore.InfoLevel
		case "warn":
			zapLevel = zapcore.WarnLevel
		case "error":
			zapLevel = zapcore.ErrorLevel
		case "fatal":
			zapLevel = zapcore.FatalLevel
		}
		// lumberjack.Logger is already safe for concurrent use, so we don't need to
		// lock it.
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename: path.Join(settings.Config.App.Log.File.Path,
				"application.logging"),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		})
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zapLevel,
		)
		logger := zap.New(core)
		logFile = logger.Sugar()

	}

}

func Debug(args ...interface{}) {
	logConsole.Debug(args...)
	if enableFileLog {
		logFile.Debug(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	logConsole.Debugf(template, args...)
	if enableFileLog {
		logFile.Debugf(template, args...)
	}
}

func Info(args ...interface{}) {
	logConsole.Info(args...)
	if enableFileLog {
		logFile.Info(args...)
	}
}

func Infof(template string, args ...interface{}) {
	logConsole.Infof(template, args...)
	if enableFileLog {
		logFile.Infof(template, args...)
	}
}

func Warn(args ...interface{}) {
	logConsole.Warn(args...)
	if enableFileLog {
		logFile.Warn(args...)
	}
}

func Warnf(template string, args ...interface{}) {
	logConsole.Warnf(template, args...)
	if enableFileLog {
		logFile.Warnf(template, args...)
	}
}

func Error(args ...interface{}) {
	logConsole.Error(args...)
	if enableFileLog {
		logFile.Error(args...)
	}
}

func Errorf(template string, args ...interface{}) {
	logConsole.Errorf(template, args...)
	if enableFileLog {
		logFile.Errorf(template, args...)
	}
}

func Fatal(args ...interface{}) {
	logConsole.Fatal(args...)
	if enableFileLog {
		logFile.Fatal(args...)
	}
}

func Fatalf(template string, args ...interface{}) {
	logConsole.Fatalf(template, args...)
	if enableFileLog {
		logFile.Fatalf(template, args...)
	}
}

func Panic(args ...interface{}) {
	logConsole.Panic(args...)
	if enableFileLog {
		logFile.Panic(args...)
	}
}

func Panicf(template string, args ...interface{}) {
	logConsole.Panicf(template, args...)
	if enableFileLog {
		logFile.Panicf(template, args...)
	}
}
