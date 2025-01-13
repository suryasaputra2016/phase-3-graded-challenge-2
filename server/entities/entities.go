package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// uers entity
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserName     string
	PasswordHash string
}

// Book entity
type Book struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string
	Author        string
	PublishedDate time.Time
	Status        string
	UserID        primitive.ObjectID `bson:"_user_id,omitempty"`
}

// Borrowed Book entity
type BorrowedBook struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	BookID       primitive.ObjectID `bson:"_book_id,omitempty"`
	UserID       primitive.ObjectID `bson:"_user_id,omitempty"`
	BorrowedDate time.Time
	ReturnDate   time.Time
}
