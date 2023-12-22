package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Topic   string             `json:"topic,omitempty"`
	Content string             `json:"content,omitempty"`
}
