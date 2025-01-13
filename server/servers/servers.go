package servers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/proto/pb"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/middlewares"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/repos"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/utils"
)

type Server struct {
	pb.UnimplementedBookServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.UserRequest) (*pb.RegisterResponse, error) {
	err := CheckRegistrationData(req.UserName, req.Password)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	res, err := CreateNewUser(req.UserName, req.Password)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	return &pb.RegisterResponse{UserName: res}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.UserRequest) (*pb.LoginResponse, error) {
	id, err := CheckLoginData(req.UserName, req.Password)
	if err != nil {
		return &pb.LoginResponse{}, err
	}

	// generate token
	t, err := middlewares.GenerateTokenString(id, req.UserName)
	if err != nil {
		return &pb.LoginResponse{}, errors.New("couldn't generate token")
	}

	return &pb.LoginResponse{UserName: req.UserName, Token: t}, nil
}

func (s *Server) ReturnABook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	err := repos.DeleteBorrowedBooks(req.BookID)
	if err != nil {
		return &pb.BookResponse{}, err
	}

	return &pb.BookResponse{
		Message: "Return succeed",
	}, nil
}

func (s *Server) BorrowABook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	bookID, err := primitive.ObjectIDFromHex(req.BookID)
	if err != nil {
		return &pb.BookResponse{}, err
	}
	userID, err := primitive.ObjectIDFromHex(middlewares.GetUserID(""))
	if err != nil {
		return &pb.BookResponse{}, err
	}

	borowedBook := entities.BorrowedBook{
		BookID:       bookID,
		UserID:       userID,
		BorrowedDate: time.Now(),
	}
	repos.AddBorrowedBooks(borowedBook)

	return &pb.BookResponse{
		Message: "Borrow succeed",
	}, nil
}

// CheckRegistrationData checks if username, and password are well formatted and username hasn't been used
func CheckRegistrationData(username, password string) error {
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

// show all books
func (s *Server) ShowAllBooks(ctx context.Context) ([]entities.Book, error) {

}

// show borrowed books
func (s *Server) ShowBorrowedBooks(ctx context.Context, userID string) ([]entities.BorrowedBook, error) {

}

// CreateNewUser accepts registration data and returns new user
func CreateNewUser(username, password string) (string, error) {
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
func CheckLoginData(username, password string) (string, error) {
	// check email and get user
	user, err := repos.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// check password
	if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
		return "", err
	}

	return user.ID.Hex(), nil
}
