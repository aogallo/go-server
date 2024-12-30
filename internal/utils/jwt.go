package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aogallo/go-server/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(payload models.User) (string, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	if JWT_SECRET == "" {
		return "", errors.New("Invalided JWT Configuration")
	}

	jwtKey := []byte(JWT_SECRET)

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		UserId:   payload.ID,
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)

}

func VerifyToken(tokenString string) (Claims, error) {

	if tokenString == "" {
		return Claims{}, errors.New("Invalided Token")
	}

	secret := getJWT()

	token, error := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}

		return secret, nil
	})

	if error != nil {
		return Claims{}, error
	}

	claims, ok := token.Claims.(Claims)

	if !ok {
		return Claims{}, errors.New("Invalid token")
	}

	if !token.Valid {
		return Claims{}, errors.New("Invalid token")
	}

	return claims, nil
}

func getJWT() []byte {

	JWT_SECRET := os.Getenv("JWT_SECRET")

	if JWT_SECRET == "" {
		return []byte("")
	}

	return []byte(JWT_SECRET)

}
