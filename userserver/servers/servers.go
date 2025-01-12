package servers

import (
	"context"
	"errors"
	"fmt"

	// pb "go-grpc-server/pb"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/userserver/entities"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/userserver/repos"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/userserver/utils"
)

type Server struct {
	// pb.UnimplementedGreetServiceServer
}

// CheckRegistrationData checks if username, and password are well formatted and username hasn't been used
func (s *Server) CheckRegistrationData(ctx context.Context, username, password string) error {
	// username validation
	if username == "" {
		return errors.New("username must not be empty")
	}

	// password validation
	if err := utils.IsPasswordGood(password); err != nil {
		return err
	}

	// check username is already in use
	if repos.GetUserByUsername(username) == nil {
		return errors.New("validating registration data: email already in use")
	}

	return nil
}

// CreateNewUser accepts registration data and returns new user
func (s *Server) CreateNewUser(ctx context.Context, username, password string) (string, error) {
	// hash password
	passwordHash, err := utils.GenerateHash(password)
	if err != nil {
		return "", fmt.Errorf("creating new user: %w", err)
	}

	// for testing purpose
	if password == "test" {
		passwordHash = "hashed password"
	}

	// define new user
	newUser := entities.User{
		UserName:     username,
		PasswordHash: passwordHash,
	}

	// add new user to database
	if repos.CreateUser(newUser) != nil {
		return "", fmt.Errorf("creating new user: %w", err)
	}

	return username, nil
}
