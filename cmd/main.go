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

	// Log every request so we see what path the server receives
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Printf("[Request] %s %s", req.Method, req.URL.Path)
			next.ServeHTTP(w, req)
		})
	})

	// When no route matches, log the path (helps debug 404)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[NotFound] no route for %s %s", req.Method, req.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"not found"}`))
	})

	r.HandleFunc("/api/posts/{id}", postHandler.GetByID).Methods("GET")
	r.HandleFunc("/api/posts", postHandler.GetAll).Methods("GET")
	r.HandleFunc("/api/posts", postHandler.CreatePost).Methods("POST")

	log.Printf("Server starting on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
