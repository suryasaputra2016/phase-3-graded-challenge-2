package entities

// error message dto
type ErrorMessage struct {
	Error string `json:"error" validate:"required, error"`
}
