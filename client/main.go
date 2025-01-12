package main

import (
	"github.com/labstack/echo/v4"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/client/handlers"
)

// @title Library
// @version 1.0
// @description Hacktiv8 Phase 3 Graded challenge 2
// @termsOfService http://swagger.io/terms/

// @contact.name Surya Saputra
// @contact.email sayasuryasaputra@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	// token := "Bearer 1234"

	userHandler := handlers.NewUserHandler()
	bookHandler := handlers.NewBookHandler()

	e.POST("/users/register", userHandler.Register)
	e.POST("/users/login", userHandler.Login)
	e.GET("/books/all", bookHandler.ShowAllBooks)
	e.GET("/books/borrowed", bookHandler.ShowBorrowedBooks)
	e.POST("/books/borrow", bookHandler.BorrowABook)
	e.PUT("/books/return", bookHandler.ReturnABook)

	e.Logger.Fatal(e.Start(":8080"))
}
