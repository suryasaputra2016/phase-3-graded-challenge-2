package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/client/entities"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/client/proto/pb"
	"google.golang.org/grpc"
)

// book handler interface
type BookHandler interface {
	BorrowABook(echo.Context) error
	ReturnABook(echo.Context) error
	ShowAllBooks(echo.Context) error
	ShowBorrowedBooks(echo.Context) error
}

// boook handler implementation
type bookHandler struct {
}

// NewBookHandler returns new book handler
func NewBookHandler() *bookHandler {
	return &bookHandler{}
}

// @Summary Borrow A Book
// @Description Borrow One Book
// @Tags books
// @Accept json
// @Produce json
// @Param borrow-book-data body entities.BorrowBookRequest true "borrowbook request"
// @Security JWT
// @Success 200 {object} entities.BorrowBookResponse
// @Router /books/borrow [post]
// @Failure 400 {object} entities.ErrorMessage
// @Failure 500 {object}  entities.ErrorMessage
func (bh *bookHandler) BorrowABook(c echo.Context) error {
	// bind request body
	var req entities.BorrowBookRequest
	if c.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.BorrowABook(ctx, &pb.BookRequest{
		BookID: req.BookID,
	})
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	return c.JSON(http.StatusOK, res)
}

// @Summary Return A Book
// @Description Return One Book
// @Tags books
// @Accept json
// @Produce json
// @Param return-book-data body entities.ReturnBookRequest true "returnbook request"
// @Security JWT
// @Success 200 {object} entities.ReturnBookResponse
// @Router /books/return [put]
// @Failure 400 {object} entities.ErrorMessage
// @Failure 500 {object}  entities.ErrorMessage
func (bh *bookHandler) ReturnABook(c echo.Context) error {
	// bind request body
	var req entities.ReturnBookRequest
	if c.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "JSON request is invalid")
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ReturnABook(ctx, &pb.BookRequest{
		BookID: req.BookID,
	})
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	return c.JSON(http.StatusOK, res)
}

// @Summary Show All Books
// @Description Show All Books in the library
// @Tags books
// @Produce json
// @Success 200 {object} []entities.ShowAllBooksResponse
// @Router /books [get]
// @Failure 500 {object}  entities.ErrorMessage
func (bh *bookHandler) ShowAllBooks(c echo.Context) error {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ShowAllBook(ctx, &pb.BookRequest{
		BookID: "",
	})
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Show Borrowed Books
// @Description Show Borrowed Books in the library
// @Tags books
// @Produce json
// @Success 200 {object} []entities.ShowBorrowedBooksResponse
// @Router /books [get]
// @Failure 500 {object}  entities.ErrorMessage
func (bh *bookHandler) ShowBorrowedBooks(c echo.Context) error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ShowBorrowedBook(ctx, &pb.BorrowedBookRequest{
		UserID: "",
	})
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	return c.JSON(http.StatusOK, res)
}
