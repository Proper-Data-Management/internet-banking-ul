package logger

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var WorkLogger, SqlLogger *zap.Logger

func WorkLoggerWithContext(ctx context.Context) *zap.Logger {
	return WithContext(WorkLogger, ctx)
}

func SqlLoggerWithContext(ctx context.Context) *zap.Logger {
	return WithContext(SqlLogger, ctx)
}

func WithContext(logger *zap.Logger, ctx context.Context) *zap.Logger {
	if metadata, ok := getMetadata(ctx); ok {
		return logger.With(zap.Any("tracing_metadata", metadata))
	}
	return logger
}

// CopyContextHolder copies ContextHolder *sync.Map from src to dst
func CopyContextHolder(dst, src context.Context) context.Context {
	return context.WithValue(dst, "ContextHolder", src.Value("ContextHolder"))
}

func getMetadata(ctx context.Context) (Metadata, bool) {
	if ctx == nil {
		return Metadata{}, false
	}
	if contextHolder, ok := ctx.Value("ContextHolder").(*sync.Map); ok {
		if value, ok := contextHolder.Load("tracing_metadata"); value != nil && ok {
			if result, ok := value.(Metadata); ok {
				return result, true
			}
		}
	}
	return Metadata{}, false
}

type Config struct {
	EnableConsole bool
	EnableFile    bool

	ConsoleJSONFormat bool
	ConsoleLevel      string

	FileJSONFormat bool
	FileLevel      string
	FileLocation   string
	FileMaxSize    int  // file size
	FileMaxBackups int  // file back
	FileMaxAge     int  // file save days
	FileCompress   bool // compress file
}

func Init(isDevelopment bool, cfg *Config) {
	if cfg == nil {
		return
	}
	WorkLogger = initLogger(isDevelopment, cfg)
	SqlLogger = initLogger(isDevelopment, cfg)
}

func initLogger(isDevelopment bool, cfg *Config) *zap.Logger {
	var cores []zapcore.Core

	if cfg.EnableConsole {
		level := getZapLevel(cfg.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(isDevelopment, cfg.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if cfg.EnableFile {
		level := getZapLevel(cfg.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: cfg.FileLocation,
			MaxSize:  cfg.FileMaxSize,
			Compress: cfg.FileCompress,
			MaxAge:   cfg.FileMaxAge,
		})
		core := zapcore.NewCore(getEncoder(isDevelopment, cfg.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	return logger
}

func getZapLevel(level string) zapcore.Level {
	var l zapcore.Level
	err := l.Set(level)
	if err != nil {
		l = zapcore.InfoLevel
	}
	return l
}

func getEncoder(isDevelopment bool, isJSON bool) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if isDevelopment {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
