package logger

import (
	"errors"
	"fmt"
	"os"

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
		panic("Imposible iniciar logger: " + err.Error())
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
	output := os.Getenv("SALIDA_LOG")
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

//Error Log error abstrae la utilidad de log implementada
func Error(msg string, err error, v ...interface{}) {
	if err == nil {
		err = errors.New("")
	}
	if len(v) > 0 {
		loggerImpl.log.Error(msg, zap.NamedError("error", err), zap.String("param", fmt.Sprintf("%v", v)))
	} else {
		loggerImpl.log.Error(msg, zap.NamedError("error", err))
	}
	loggerImpl.log.Sync()
}
