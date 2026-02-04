package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dhyaneshsiddhartha15/crud-go/internal/model"
	"github.com/dhyaneshsiddhartha15/crud-go/internal/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req model.CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	authorID := "507f1f77bcf86cd799439011"

	if err := h.service.CreatePost(r.Context(), &req, authorID); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := strings.TrimSpace(rawID)
	log.Printf("[GetByID handler] raw id=%q len=%d, trimmed id=%q len=%d", rawID, len(rawID), id, len(id))

	if id == "" {
		http.Error(w, "Post ID required", http.StatusBadRequest)
		return
	}
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		log.Printf("[GetByID handler] invalid ObjectID: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid post ID"})
		return
	}

	post, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		log.Printf("[GetByID handler] service error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Post not found"})
		return
	}
	log.Printf("[GetByID handler] found post id=%s", post.ID.Hex())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
