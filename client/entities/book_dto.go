package entities

// borrow book request dto
type BorrowBookRequest struct {
	BookID string `json:"book_id" validate:"required, book_id"`
}

// borrow book response dto
type BorrowBookResponse struct {
	Message    string `json:"message" validate:"required, message"`
	BookTitle  string `json:"book_title" validate:"required, book_title"`
	BookAuthor string `json:"book_author" validate:"required, book_author"`
}

// return book request dto
type ReturnBookRequest struct {
	BookID string `json:"book_id" validate:"required, book_id"`
}

// return book response dto
type ReturnBookResponse struct {
	Message    string `json:"message" validate:"required, message"`
	BookTitle  string `json:"title" validate:"required, title"`
	BookAuthor string `json:"author" validate:"required, author"`
}

// show all book response dto
type ShowAllBooksResponse struct {
	BookID        string `json:"book_id" validate:"required, book_id"`
	Title         string `json:"title" validate:"required, title"`
	Author        string `json:"author" validate:"required, author"`
	PublishedDate string `json:"published_date" validate:"required, published_date"`
	Status        string `json:"status" validate:"required, status"`
}

// show borrowed book response dto
type ShowBorrowedBooksResponse struct {
	BookID       string `json:"book_id" validate:"required, book_id"`
	BorrowedDate string `json:"published_date" validate:"required, published_date"`
}
