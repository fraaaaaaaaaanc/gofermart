package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	Log  *zap.Logger
	File *os.File
}

var l Logger

func NewZapLogger(filePath, logLvl string) error {
	var core zapcore.Core
	var file *os.File

	switch logLvl {
	case envLocal:
		consoleConfig := zap.NewDevelopmentConfig()
		consoleConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		core = zapcore.NewCore(zapcore.NewConsoleEncoder(consoleConfig.EncoderConfig),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel)
	case envDev:
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		writeSyncer := zapcore.AddSync(file)
		fileConfig := zap.NewProductionConfig()
		fileConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		core = zapcore.NewCore(zapcore.NewJSONEncoder(fileConfig.EncoderConfig),
			writeSyncer,
			zapcore.InfoLevel)
	}
	log := zap.New(core, zap.AddCaller())
	l = Logger{File: file, Log: log}
	return nil
}

func CloseFile() {
	if l.File != nil {
		if err := l.File.Close(); err != nil {
			l.Log.Error("Failed to close log file", zap.Error(err))
		}
	}
}

// Info is a method to log informational messages.
func Info(msg string, fields ...zapcore.Field) {
	l.Log.Info(msg, fields...)
}

// Error is a method to log error messages.
func Error(msg string, fields ...zapcore.Field) {
	l.Log.Error(msg, fields...)
}

func With(userID int, err error, r *http.Request) {
	l.Log.Info("Request details",
		zap.Int("UserID", userID),
		zap.String("uri", r.RequestURI),
		zap.String("method", r.Method),
		zap.Error(err),
	)
}
