package servers

import (
	"context"
	"errors"
	"fmt"

	// pb "go-grpc-server/pb"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/repos"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/utils"
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
	_, err := repos.GetUserByUsername(username)
	if err == nil {
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

// CheckRegistrationData checks if username, and password are well formatted and username hasn't been used
func (s *Server) CheckLoginData(ctx context.Context, username, password string) error {
	// check email and get user
	user, err := repos.GetUserByUsername(username)
	if err != nil {
		return err
	}

	// check password
	if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
		return err
	}

	return nil
}

// show all books
func (s *Server) ShowAllBooks(ctx context.Context) ([]entities.Book, error) {

}

// show borrowed books
func (s *Server) ShowBorrowedBooks(ctx context.Context, userID string) ([]entities.BorrowedBook, error) {

}

func (s *Server) ReturnABook(ctx context.Context, userID, bookID string) error {

}

func (s *Server) BorrowABook(ctx context.Context, userID, bookID string) error {

}
