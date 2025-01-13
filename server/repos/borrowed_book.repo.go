package repos

import (
	"context"
	"errors"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/configs"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get borrowed books by user ID
func GetBorrowedBooks(borrowerID string) ([]entities.BorrowedBook, error) {
	userID, err := primitive.ObjectIDFromHex(borrowerID)
	if err != nil {
		return []entities.BorrowedBook{}, err
	}

	collection, err := configs.BorrowedBookDatabase(context.Background())
	if err != nil {
		return []entities.BorrowedBook{}, err
	}

	var datas []entities.BorrowedBook

	// Proses get all data
	cursor, err := collection.Find(context.Background(), bson.M{"_user_id": userID})
	if err != nil {
		return []entities.BorrowedBook{}, err
	}

	// Looping Cursor
	for cursor.Next(context.Background()) {
		var data entities.BorrowedBook
		if err := cursor.Decode(&data); err != nil {
			return []entities.BorrowedBook{}, err
		}

		// proses append "data" ke dalam "datas"
		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		return []entities.BorrowedBook{}, err
	}

	return datas, nil
}

// add borrowed books
func AddBorrowedBooks(borrowedBook entities.BorrowedBook) error {
	collection, err := configs.BorrowedBookDatabase(context.Background())
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), borrowedBook)
	if err != nil {
		return err
	}

	return err
}

// delete borrowed books
func DeleteBorrowedBooks(bookID string) error {
	collection, err := configs.BorrowedBookDatabase(context.Background())
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return err
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("borrowed book not found")
	}

	return nil
}
