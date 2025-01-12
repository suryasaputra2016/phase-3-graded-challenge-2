package configs

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection Database connect to remote mongodb user collection
func UserDatabase(ctx context.Context) (*mongo.Collection, error) {
	connection := "mongodb+srv://Cluster11072:aUhQSlddVlRv@cluster11072.y8sq9.mongodb.net/?retryWrites=true&w=majority&appName=Cluster11072"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	collection := client.Database("bookborrow").Collection("users")
	return collection, nil
}

// Connection Database connect to remote mongodb book collection
func BookDatabase(ctx context.Context) (*mongo.Collection, error) {
	connection := "mongodb+srv://Cluster11072:aUhQSlddVlRv@cluster11072.y8sq9.mongodb.net/?retryWrites=true&w=majority&appName=Cluster11072"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	collection := client.Database("bookborrow").Collection("books")
	return collection, nil
}

func BorrowedBookDatabase(ctx context.Context) (*mongo.Collection, error) {
	connection := "mongodb+srv://Cluster11072:aUhQSlddVlRv@cluster11072.y8sq9.mongodb.net/?retryWrites=true&w=majority&appName=Cluster11072"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	collection := client.Database("bookborrow").Collection("Borrowedbooks")
	return collection, nil
}
