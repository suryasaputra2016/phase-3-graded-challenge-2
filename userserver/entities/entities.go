package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// uers entity
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserName     string
	PasswordHash string
}
