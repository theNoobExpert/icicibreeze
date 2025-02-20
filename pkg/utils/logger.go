package utils

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

// ParseLogLevel converts string log level to zapcore.Level. Defaults to 'debug'.
func ParseLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func init() {
	logLevel := os.Getenv("BREEZE_LOG_LEVEL")
	if logLevel == "" {
		logLevel = "debug"
	}

	logFilePath := os.Getenv("BREEZE_LOG_FILE_PATH")
	level := ParseLogLevel(logLevel)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeTime:   zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Console core (stdout logging)
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		level,
	)

	var cores []zapcore.Core
	cores = append(cores, consoleCore)

	// File rotating core (if log file is provided)
	if logFilePath != "" {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    10, // MB before rotation
			MaxBackups: 5,  // Max number of old log files
			MaxAge:     30, // Days to retain
			Compress:   true,
		}
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(lumberjackLogger),
			level,
		)
		cores = append(cores, fileCore)
	}

	core := zapcore.NewTee(cores...)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	logger = zapLogger.Sugar()

	// Ensure logs are flushed
	defer zapLogger.Sync()

	logger.Infof("Logger initialized (Level: %s, Log file: %s)", logLevel, logFilePath)
}

// GetLogger returns the global logger instance.
func GetLogger() *zap.SugaredLogger {
	return logger
}
