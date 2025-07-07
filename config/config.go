package config

import (
	"fmt"
	"os"

	"go-commerce/schemas"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	logger = NewLogger("config")

	dbPath := "./db/main.db"

	if err := os.MkdirAll("./db", os.ModePerm); err != nil {
		return fmt.Errorf("error creating db directory: %v", err)
	}
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		logger.Info("database file not found, creating ...")
		f, err := os.Create(dbPath)
		if err != nil {
			return fmt.Errorf("error creating db file: %v", err)
		}
		f.Close()
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return fmt.Errorf("error initializing sqlite: %w", err)
	}

	if err := db.AutoMigrate(
		&schemas.User{},
		&schemas.Product{},
		&schemas.Sale{},
	); err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return fmt.Errorf("error running automigrate: %w", err)
	}

	logger.Info("sqlite initialized and migrated successfully")
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	if logger == nil {
		logger = NewLogger(prefix)
	}
	return logger
}
