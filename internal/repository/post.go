package repository

import (
	"context"
	"log"

	"github.com/dhyaneshsiddhartha15/crud-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	for i, p := range posts {
		log.Printf("[GetAll repo] post[%d] _id=%s (hex)", i, p.ID.Hex())
	}
	return posts, nil
}

func (r *PostRepository) GetByID(ctx context.Context, postID string) (*model.Post, error) {
	objectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Printf("[GetByID repo] ObjectIDFromHex failed: postID=%q err=%v", postID, err)
		return nil, err
	}
	log.Printf("[GetByID repo] querying _id=%s (hex)", objectID.Hex())

	var post model.Post
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		log.Printf("[GetByID repo] FindOne failed: err=%v", err)
		return nil, err
	}
	log.Printf("[GetByID repo] found post _id=%s", post.ID.Hex())
	return &post, nil
}
