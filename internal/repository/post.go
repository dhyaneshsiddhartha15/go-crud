package repository

import (
	"context"

	"github.com/dhyaneshsiddhartha15/crud-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{
		collection: db.Collection("posts"),
	}
}
func (r *PostRepository) Create(ctx context.Context, post *model.Post) error {
	_, err := r.collection.InsertOne(ctx, post)
	return err
}
func (r *PostRepository) GetAll(ctx context.Context) ([]model.Post, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []model.Post

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}
