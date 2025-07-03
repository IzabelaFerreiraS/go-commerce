package main

import (
	"github.com/IzabelaFerreiraS/go-commerce.git/config"
	"github.com/IzabelaFerreiraS/go-commerce.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}