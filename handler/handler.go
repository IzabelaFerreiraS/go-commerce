package handler

import (
	"github.com/IzabelaFerreiraS/go-commerce.git/config"
	"gorm.io/gorm"
)

var (
	logger  *config.Logger
	db      *gorm.DB
)

func InitializeHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()

	logger.Info("Handler initialized")
}
