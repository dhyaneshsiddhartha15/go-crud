package services

import (
	"context"
	"time"

	"github.com/yourusername/crud-go/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(
	ctx context.Context,
	collection *mongo.Collection,
	title string,
	description  string,
	authorID primitive.ObjectID,
) (*models.Post,error){ 
	newPost:=&models.Post{
		ID:primitive.NewObjectID(),
		Title:title,
		Description:description,
		AuthorID:authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
_,err:=collection.InsertOne(ctx,newPost)
if err!=nil{
	return nil,err
}
return  newPost,nil

	}
