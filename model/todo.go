package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id      primitive.ObjectID `json:"_id,omitempty" `
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
}
