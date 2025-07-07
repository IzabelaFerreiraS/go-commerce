package main

import (
	"go-commerce/config"
	"go-commerce/routes"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime)

	if err := config.Init(); err != nil {
		logger.Printf("config initialization error: %v", err)
		return
	}

	logger.Println("configuration completed, starting server")
	routes.Initialize()
}
