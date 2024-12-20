package config

import "github.com/rs/zerolog"

type LoggerConfig struct {
	Level      zerolog.Level
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SslMode  string
}

type ServerConfig struct {
	Port int
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Port: 8080,
	}
}

func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Level:      zerolog.InfoLevel,
		FilePath:   "../../logs/global.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
}

func DefaultDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "password",
		Name:     "auth_example",
		SslMode:  "disable",
	}
}
