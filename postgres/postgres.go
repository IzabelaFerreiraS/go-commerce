package postgres

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func New(dsn string, gormConfig *gorm.Config) *Postgres {
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)

	if err != nil {
		return nil
	}

	return &Postgres{db}
}

func (p *Postgres) RunPing(maxIdleConns int, maxOpenconns int, connMaxLifetime time.Duration) error {
	sqlDB, err := p.DB.DB()

	if err != nil {
		log.Fatalf("Falha ao configurar pool de conexões: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenconns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	err = sqlDB.Ping()

	if err != nil {
		log.Fatalf("Falha ao fazer ping no banco de dados: %v", err)
		return err
	}

	return nil
}

func (p *Postgres) RunAutoMigrate(models ...any) error {
	if err := p.DB.AutoMigrate(
		models...,
	); err != nil {
		log.Fatalf("Falha no auto migrate: %v", err)
		return err
	}

	return nil
}

func (p *Postgres) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
