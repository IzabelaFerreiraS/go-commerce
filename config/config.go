package config

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return errors.New("fake error")
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger{
	logger = NewLogger(p)
	return logger
}