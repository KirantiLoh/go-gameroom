package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserData
}

func EncodeJWT(user UserData) (string, error) {
	claims := &Claims{
		UserData: user,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   string(user.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // specify duration clearly
		},
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func DecodeJWT(tokenString string) (*Claims, error) {

	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return nil, errors.New("JWT_SECRET environment variable is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil // returning the secret as []byte
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
