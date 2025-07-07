package postgres

import (
	"log"
	"os"

	"gorm.io/gorm/logger"
)

type PostgresLogger struct {
	Instance logger.Interface
}

func NewLogger(config logger.Config) *PostgresLogger {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		config,
	)

	return &PostgresLogger{newLogger}
}

func (l *PostgresLogger) GetLogLevel() logger.LogLevel {
	return logger.Info
}
