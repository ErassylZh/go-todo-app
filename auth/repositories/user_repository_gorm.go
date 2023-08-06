package repositories

import (
	"gorm.io/gorm"
	"my_first_go_project/auth/models"
)

type UserRepository interface {
	CreateUser(user models.User)
	GetUserByUsername(username string) (models.User, error)
}

type UserRepositoryGORM struct {
	db *gorm.DB
}

func NewUserRepositoryGORM(db *gorm.DB) *UserRepositoryGORM {
	return &UserRepositoryGORM{db: db}
}

func (ur *UserRepositoryGORM) CreateUser(user models.User) {
	ur.db.Create(&user)
}

func (ur *UserRepositoryGORM) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	result := ur.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
