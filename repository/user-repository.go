package repository

import (
	"github.com/yet-another-todo-list-golang/config"
	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (error, entity.User)
	Create(user entity.User) (error, entity.User)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db := config.DatabaseConnect()
	return &userRepository{
		db: db,
	}
}

func (userRepository *userRepository) FindByEmail(email string) (error, entity.User) {
	var user entity.User

	if err := userRepository.db.First(&user, "email = ?", email).Error; err != nil {
		return err, user
	}

	return nil, user
}

func (userRepository *userRepository) Create(user entity.User) (error, entity.User) {
	if err := userRepository.db.Create(&user).Error; err != nil {
		return err, user
	}
	return nil, user
}
