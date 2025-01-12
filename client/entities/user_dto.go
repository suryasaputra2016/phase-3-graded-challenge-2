package entities

// reqister request dto
type RegisterRequest struct {
	UserName string `json:"user_name" validate:"required, user_name"`
	Password string `json:"password" validate:"required, password"`
}

// reqister response dto
type RegisterResponse struct {
	Message  string `json:"message" validate:"required, message"`
	UserName string `json:"user_name" validate:"required, user_name"`
}

// login request dto
type LoginRequest struct {
	UserName string `json:"user_name" validate:"required, user_name"`
	Password string `json:"password" validate:"required, password"`
}

// login response dto
type LoginResponse struct {
	Message string `json:"message" validate:"required, message"`
	Token   string `json:"token" validate:"required, token"`
}
