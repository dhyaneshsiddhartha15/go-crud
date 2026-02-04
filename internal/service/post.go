package service

import (
	"context"
	"log"
	"time"

	"github.com/dhyaneshsiddhartha15/crud-go/internal/model"
	"github.com/dhyaneshsiddhartha15/crud-go/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(
	ctx context.Context,
	req *model.CreatePostRequest,
	authorID string,
) error {

	objID, err := primitive.ObjectIDFromHex(authorID)
	if err != nil {
		return err
	}

	post := &model.Post{
		ID:          primitive.NewObjectID(),
		Title:       req.Title,
		Description: req.Description,
		AuthorID:    objID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	log.Printf("I created this", post)

	return s.repo.Create(ctx, post)
}

func (s *PostService) GetAll(ctx context.Context) ([]model.Post, error) {
	posts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
