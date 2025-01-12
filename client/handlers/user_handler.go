package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/client/entities"
)

// user handler interface
type UserHandler interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

// user handler implementation
type userHandler struct {
}

// NewUserHandler takes user service and returns new user handler
func NewUserHandler() *userHandler {
	return &userHandler{}
}

// @Summary Register
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param register-data body entity.RegisterRequest true "register request"
// @Success 201 {object} entity.RegisterRepsonse
// @Router /register [post]
// @Failure 400 {object} entity.ErrorMessage
// @Failure 500 {object}  entity.ErrorMessage
func (uh *userHandler) Register(c echo.Context) error {
	// bind request body
	var req entities.RegisterRequest
	if c.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	// check registration data
	// if err := uh.us.CheckRegistrationData(req.Email, req.Password); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint(err))
	// }

	// create new user
	// newUserPtr, err := uh.us.CreateNewUser(&req)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprint(err))
	// }

	//define and send response
	res := entities.RegisterResponse{
		Message:  "registration success",
		UserName: "xxxxxxxxxxx",
	}
	return c.JSON(http.StatusCreated, res)
}

// @Summary Login
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param login-data body entity.LoginRequest true "login request"
// @Success 200 {object} entity.LoginResponse
// @Router /login [post]
// @Failure 400 {object} entity.ErrorMessage
// @Failure 500 {object}  entity.ErrorMessage
func (uh *userHandler) Login(c echo.Context) error {
	// bind request body
	var req entities.LoginRequest
	if c.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	// check login data
	// userPtr, err := uh.us.CheckLoginData(req.Email, req.Password)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint(err))
	// }

	// generate token
	// t, err := middlewares.GenerateTokenString(int(userPtr.ID), userPtr.Email)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "couldn't generate token")
	// }

	// define and send response
	res := entities.LoginResponse{
		Message: "login success",
		Token:   "token xxxxxxx",
	}
	return c.JSON(http.StatusOK, res)
}
