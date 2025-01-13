package repos

import (
	"context"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/configs"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
	"go.mongodb.org/mongo-driver/bson"
)

// Create User
func CreateUser(newUser entities.User) error {
	collection, err := configs.UserDatabase(context.Background())
	if err != nil {
		return err
	}
	//  insert data to mongo
	_, err = collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return err
	}

	return nil
}

// Get User by username
func GetUserByUsername(userName string) (entities.User, error) {
	collection, err := configs.UserDatabase(context.Background())
	if err != nil {
		return entities.User{}, err
	}

	// get user
	var user entities.User
	err = collection.FindOne(context.Background(), bson.M{"userName": userName}).Decode(&user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
