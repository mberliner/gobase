package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerImpl logger
)

type appLogger interface {
	Print(v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if loggerImpl.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func getLevel() zapcore.Level {
	switch os.Getenv("NIVEL_LOG") {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "ERROR":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	output := strings.TrimSpace(os.Getenv("SALIDA_LOG"))
	if output == "" {
		return "stdout"
	}
	return output
}

//GetLogger expone un logger genérico
func GetLogger() appLogger {
	return loggerImpl
}

//logger driver Mysql
func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

//Info Log info abstrae la utilidad de log implementada
func Info(msg string) {
	loggerImpl.log.Info(msg)
	loggerImpl.log.Sync()
}

//Error Log errro abstrae la utilidad de log implementada
func Error(msg string, err error) {
	loggerImpl.log.Error(msg, zap.NamedError("error", err))
	loggerImpl.log.Sync()
}
