package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	token := "Bearer 1234"

	e.POST("/users/register", handlr.Register)
	e.POST("/users/login", handlr.Login)
	e.GET("/books/all", handlr.ShowBooks)
	e.GET("/books/borrowed", handlr.ShowBorrowedBooks)
	e.POST("/books/borrow/:bookId", handlr.Borrow)
	e.PUT("/books/return/:bookId", handlr.Return)

	e.Logger.Fatal(e.Start(":8080"))
}
