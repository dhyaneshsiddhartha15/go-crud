package main

import (
	"log"

	"github.com/yourusername/crud-go/internal/config"
	"github.com/yourusername/crud-go/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.MongoURI, cfg.MongoDBName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	_ = db
}
