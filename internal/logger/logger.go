package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type ZapLogger struct {
	Log  *zap.Logger
	File *os.File
}

func NewZapLogger(filePath, logLvl string) (*ZapLogger, error) {
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
			return nil, err
		}
		writeSyncer := zapcore.AddSync(file)
		fileConfig := zap.NewProductionConfig()
		fileConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		core = zapcore.NewCore(zapcore.NewJSONEncoder(fileConfig.EncoderConfig),
			writeSyncer,
			zapcore.InfoLevel)
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &ZapLogger{File: file, Log: logger}, nil
}

func (z *ZapLogger) CloseFile() {
	if z.File != nil {
		if err := z.File.Close(); err != nil {
			z.Error("Failed to close log file", zap.Error(err))
		}
	}
}

// Info is a method to log informational messages.
func (z *ZapLogger) Info(msg string, fields ...zapcore.Field) {
	z.Log.Info(msg, fields...)
}

// Error is a method to log error messages.
func (z *ZapLogger) Error(msg string, fields ...zapcore.Field) {
	z.Log.Error(msg, fields...)
}
