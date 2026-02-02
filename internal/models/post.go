package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"` // Fixed: was unexported
	AuthorID    primitive.ObjectID `json:"author_id" bson:"author_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"` // Fixed: added json and bson tags
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
