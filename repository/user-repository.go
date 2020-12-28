package repository

import (
	"github.com/yet-another-todo-list-golang/db"
	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (error, entity.User)
	Create(user entity.User) (error, entity.User)
	FindById(id string) (error, entity.User)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	connection := db.GetConnection()
	return &userRepository{
		connection: connection,
	}
}

func (userRepository *userRepository) FindByEmail(email string) (error, entity.User) {
	var user entity.User

	if err := userRepository.connection.First(&user, "email = ?", email).Error; err != nil {
		return err, user
	}

	return nil, user
}

func (userRepository *userRepository) Create(user entity.User) (error, entity.User) {
	if err := userRepository.connection.Create(&user).Error; err != nil {
		return err, user
	}
	return nil, user
}

func (userRepository *userRepository) FindById(id string) (error, entity.User) {
	var user entity.User

	if err := userRepository.connection.First(&user, "id = ?", id).Error; err != nil {
		return err, user
	}

	return nil, user
}
