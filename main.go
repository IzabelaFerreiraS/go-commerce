package main

import (
	"go-commerce/config"
	"go-commerce/routes"
)

func main() {
    if err := config.Init(); err != nil {
        config.GetLogger("main").Errorf("config initialization error: %v", err)
        return
    }

    logger := config.GetLogger("main")
    logger.Info("configuration completed, starting server")

    routes.Initialize()
}
