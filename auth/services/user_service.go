package services

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"my_first_go_project/auth/models"
	"my_first_go_project/auth/repositories"
	"time"
)

type UserService interface {
	RegisterUser(user models.User)
	LoginUser(loginData models.LoginData) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{userRepository: ur}
}

func (us *userService) RegisterUser(user models.User) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		return
	}
	user.Password = string(hashedPassword)

	us.userRepository.CreateUser(user)
}
func (us *userService) LoginUser(loginData models.LoginData) (string, error) {
	user, err := us.userRepository.GetUserByUsername(loginData.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return "", err
	}

	token := generateAuthToken(user.ID)
	return token, nil
}

func generateAuthToken(userID uint) string {
	secretKey := []byte("super_hardest_secret_key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err) // или вернуть пустую строку вместо panic
	}

	return tokenString
}
