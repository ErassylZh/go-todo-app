package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GetUserFromToken(tokenString string) (uint, error)
}

type authService struct {
	jwtSecret []byte
}

func NewAuthService(jwtSecret string) AuthService {
	return &authService{
		jwtSecret: []byte(jwtSecret),
	}
}

// super_hardest_secret_key
func (as *authService) GetUserFromToken(tokenString string) (uint, error) {
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte("super_hardest_secret_key"), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("Invalid token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		userIDInt, ok := claims["user_id"].(int) // Или int64, в зависимости от типа в токене
		if !ok {
			return 0, errors.New("Invalid user ID in token")
		}
		userID = float64(userIDInt)
	}

	return uint(userID), nil
}
