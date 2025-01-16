package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aogallo/go-server/internal/v1/models"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
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

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &CustomClaims{
		UserId:   payload.ID,
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (models.UserResponse, error) {
	if tokenString == "" {
		return models.UserResponse{}, errors.New("Invalided Token")
	}

	secret := getJWT()

	token, error := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}

		return secret, nil
	})

	if error != nil {
		return models.UserResponse{}, errors.New(error.Error())
	} else if claims, ok := token.Claims.(*CustomClaims); ok {

		fmt.Println(&claims.Username, &claims.UserId)

		return models.UserResponse{
			ID:       claims.UserId,
			Username: claims.Username,
		}, nil
	} else {
		return models.UserResponse{}, errors.New("Invalid credentials")
	}
}

func getJWT() []byte {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	if JWT_SECRET == "" {
		return []byte("")
	}

	return []byte(JWT_SECRET)
}
