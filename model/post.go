package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string              `json:"title,omitempty" bson:"title" validate:"required"`
	Content     string              `json:"content,omitempty" bson:"content" validate:"required"`
	CreatedDate time.Time           `json:"created_date,omitempty" bson:"cteated_date"`
}
