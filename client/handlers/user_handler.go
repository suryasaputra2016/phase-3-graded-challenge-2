package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/client/entities"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/proto/pb"
	"google.golang.org/grpc"
)

// user handler interface
type UserHandler interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

// user handler implementation
type userHandler struct {
}

// NewUserHandler returns new user handler
func NewUserHandler() *userHandler {
	return &userHandler{}
}

// @Summary Register
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param register-data body entities.RegisterRequest true "register request"
// @Success 201 {object} entities.RegisterResponse
// @Router /register [post]
// @Failure 400 {object} entities.ErrorMessage
// @Failure 500 {object}  entities.ErrorMessage
func (uh *userHandler) Register(c echo.Context) error {
	// bind request body
	var req entities.RegisterRequest
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

	res, err := client.Register(ctx, &pb.UserRequest{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	//send response
	return c.JSON(http.StatusCreated, res)
}

// @Summary Login
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param login-data body entities.LoginRequest true "login request"
// @Success 200 {object} entities.LoginResponse
// @Router /login [post]
// @Failure 400 {object} entities.ErrorMessage
// @Failure 500 {object}  entities.ErrorMessage
func (uh *userHandler) Login(c echo.Context) error {
	// bind request body
	var req entities.LoginRequest
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

	res, err := client.Login(ctx, &pb.UserRequest{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	return c.JSON(http.StatusOK, res)
}
