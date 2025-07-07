package config

import (
	"fmt"
	"go-commerce/postgres"
	"go-commerce/schemas"
	"time"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AppConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`
	Port        string `env:"PORT" envDefault:"8080"`
	PostgresDsn string `env:"POSTGRES_DSN"`
}

var (
	Env *AppConfig
	DB *postgres.Postgres
)

func Load() error {
	Env = new(AppConfig)
	if err := env.Parse(Env); err != nil {
		return fmt.Errorf("failed to parse environment variables: %w", err)
	}
	return nil
}

func Init() error {
	if err := Load(); err != nil {
		return err
	}

	dbLogger := postgres.NewLogger(logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	DB = postgres.New(Env.PostgresDsn, &gorm.Config{
		Logger: dbLogger.Instance,
	})
	if DB == nil {
		return fmt.Errorf("failed to initialize postgres connection")
	}

	if err := DB.RunPing(10, 100, time.Minute); err != nil {
		return err
	}

	if err := DB.RunAutoMigrate(
		&schemas.User{},
		&schemas.Product{},
		&schemas.Sale{},
	); err != nil {
		return err
	}

	return nil
}

