package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/client/entities"
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

	// check book rental requirements
	// userID := middlewares.GetUserID(c.Get("user"))
	// bookPtr, userPtr, err := bh.bs.CheckBookRentalRequirements(req.Title, req.Author, userID)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint(err))
	// }

	// process book rental
	// newRent, err := bh.bs.ProcessBookRental(bookPtr, userPtr)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprint(err))
	// }

	// define and send response
	res := entities.BorrowBookResponse{
		Message:    "borrow success",
		BookTitle:  "xxxx",
		BookAuthor: "xxxxxxx",
	}
	return c.JSON(http.StatusOK, res)
}

// @Summary Return A Book
// @Description Return One Book
// @Tags books
// @Accept json
// @Produce json
// @Param return-book-data body entites.ReturnBookRequest true "returnbook request"
// @Security JWT
// @Success 200 {object} entities.ReturnBookResponse
// @Router /books/return [put]
// @Failure 400 {object} entity.ErrorMessage
// @Failure 500 {object}  entity.ErrorMessage
func (bh *bookHandler) ReturnABook(c echo.Context) error {
	// bind request body
	var req entities.ReturnBookRequest
	if c.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "JSON request is invalid")
	}

	// check book return requirements
	// userID := middlewares.GetUserID(c.Get("user"))
	// rentPtr, err := bh.bs.CheckBookReturnRequirements(req.Title, req.Author, req.CopyNumber, userID)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint(err))
	// }

	// process book return
	// copyPtr, err := bh.bs.ProcessBookReturn(rentPtr)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprint(err))
	// }

	// record in rent history
	// if err = bh.bs.StoreRentHistory(uint(userID), copyPtr.ID, "return"); err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprint(err))
	// }

	// define and send response
	res := entities.ReturnBookResponse{
		Message:    "return success",
		BookTitle:  "Book xxx",
		BookAuthor: "Book xxx",
	}
	return c.JSON(http.StatusOK, res)
}

// @Summary Show All Books
// @Description Show All Books in the library
// @Tags books
// @Produce json
// @Success 200 {object} entities.ShowAllBooksResponse
// @Router /books [get]
// @Failure 500 {object}  entity.ErrorMessage
func (bh *bookHandler) ShowAllBooks(c echo.Context) error {
	// get all books
	// bookCopiesPtr, err := bh.bs.GetAllBooks()
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprint(err))
	// }

	// define and send response
	var res []entities.ShowAllBooksResponse
	var copyResponse entities.ShowAllBooksResponse
	for _, copy := range []string{"xxxxx"} { //*bookCopiesPtr {
		copyResponse = entities.ShowAllBooksResponse{
			BookID:        copy,
			Title:         "",
			Author:        "",
			PublishedDate: "",
			Status:        "",
		}
		res = append(res, copyResponse)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Show Borrowed Books
// @Description Show Borrowed Books in the library
// @Tags books
// @Produce json
// @Success 200 {object} entities.ShowBorrowedBooksResponse
// @Router /books [get]
// @Failure 500 {object}  entity.ErrorMessage
func (bh *bookHandler) ShowBorrowedBooks(c echo.Context) error {
	// get all books
	// bookCopiesPtr, err := bh.bs.GetAllBooks()
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprint(err))
	// }

	// define and send response
	var res []entities.ShowBorrowedBooksResponse
	var copyResponse entities.ShowBorrowedBooksResponse
	for _, copy := range []string{"xxxxx"} { //*bookCopiesPtr {
		copyResponse = entities.ShowBorrowedBooksResponse{
			BookID:       copy,
			BorrowedDate: "",
		}
		res = append(res, copyResponse)
	}

	return c.JSON(http.StatusOK, res)
}
