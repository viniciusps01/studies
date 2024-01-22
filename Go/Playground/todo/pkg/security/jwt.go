package security

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/viniciusps01/todo/pkg/apperrors"
)

const jwtKey = "JWT_KEY"

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var secret []byte

func init() {
	godotenv.Load()

	k := os.Getenv(jwtKey)

	if k == "" {
		log.Fatal("JWT key can not be EMPTY")
	}

	secret = []byte(k)
}

func GenerateToken(c Claims) (*string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	s, err := t.SignedString(secret)

	if err != nil {
		return nil, apperrors.InternalServerError{
			Message: "failed to generate JWT token: " + err.Error(),
		}
	}

	return &s, nil
}

func ValidateToken(token string) (*Claims, error) {
	var claims Claims

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		switch t.Method.(type) {
		case *jwt.SigningMethodHMAC:
			return secret, nil

		default:
			return nil, apperrors.BadRequestError{
				Message: "invalid jwt token",
			}
		}
	})

	if err != nil || !parsedToken.Valid {
		return nil, apperrors.BadRequestError{
			Message: "invalid jwt token",
		}
	}

	return &claims, nil
}
