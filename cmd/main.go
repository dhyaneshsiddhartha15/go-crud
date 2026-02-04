package main

import (
	"log"
	"net/http"

	"github.com/dhyaneshsiddhartha15/crud-go/internal/config"
	"github.com/dhyaneshsiddhartha15/crud-go/internal/database"
	"github.com/dhyaneshsiddhartha15/crud-go/internal/handler"
	"github.com/dhyaneshsiddhartha15/crud-go/internal/repository"
	"github.com/dhyaneshsiddhartha15/crud-go/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.MongoURI, cfg.MongoDBName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", postHandler.GetAll).Methods("GET")
	r.HandleFunc("/api/posts", postHandler.CreatePost).Methods("POST")

	log.Printf("Server starting on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
