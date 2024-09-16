package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	Evn string
}

func (Log Config) selectLogLevel() zap.AtomicLevel {
	switch Log.Evn {
	case "local":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "dev":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "prod":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func Setup(Log Config) *zap.Logger {
	level := Log.selectLogLevel()

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "ts"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "tmp/todo-app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)

	log := zap.New(core, zap.AddCaller())

	return log
}
