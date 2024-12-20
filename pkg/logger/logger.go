package logger

import (
	"auth/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(config config.LoggerConfig) zerolog.Logger {
	writer := &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	multiWriter := zerolog.MultiLevelWriter(os.Stdout, writer)
	return zerolog.New(multiWriter).Level(config.Level).With().Timestamp().Logger()
}

func InitializeDistributedLoggers() {
	config := config.DefaultLoggerConfig()
	globalLogger := NewLogger(config)
	log.Logger = globalLogger
}
