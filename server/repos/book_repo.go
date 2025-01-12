package repos

import (
	"context"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/configs"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
	"go.mongodb.org/mongo-driver/bson"
)

// Get all books
func GetAllBooks(userName string) ([]entities.Book, error) {
	collection, err := configs.BookDatabase(context.Background())
	if err != nil {
		return []entities.Book{}, err
	}

	var datas []entities.Book

	// Proses get all data
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return []entities.Book{}, err
	}

	// Looping Cursor
	for cursor.Next(context.Background()) {
		var data entities.Book
		if err := cursor.Decode(&data); err != nil {
			return []entities.Book{}, err
		}

		// proses append "data" ke dalam "datas"
		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		return []entities.Book{}, err
	}

	return datas, nil
}

// Get borrowed books by user ID
func GetBorrowedBooks(userID string) ([]entities.BorrowedBook, error) {
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
