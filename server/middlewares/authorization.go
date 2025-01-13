package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// custom jwt claim
type jwtCustomClaims struct {
	ID       string `json:"id"`
	UserName string `json:"UserName"`
	jwt.RegisteredClaims
}

// GetUserEmail accepts echo context and returns email data from jwt encoded token
func GetUsername(c interface{}) string {
	user := c.(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims.UserName
}

// GetUserID accepts echo context and returns id data from jwt encoded token
func GetUserID(c interface{}) string {
	user := c.(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims.ID
}

// GenerateTokenString generates token string with custom claims and returns encoded token and error
func GenerateTokenString(id string, userName string) (string, error) {
	// Set custom claim
	claims := &jwtCustomClaims{
		id,
		userName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// Authorization returns middleware jwt authorization
func Authorization() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtCustomClaims)
			},
			SigningKey: []byte(os.Getenv("JWT_SECRET")),
		},
	)
}
