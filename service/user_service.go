package service

import (
	"errors"
	"github.com/yet-another-todo-list-golang/helper"
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/repository"
	"log"
)

type UserService interface {
	Login(loginDto dto.LoginDto) (error, string, entity.User)
	Register(registerDto dto.RegisterUserDto) (error, entity.User)
	GetOne(id string) (error, entity.User)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

//Login returns error for error happens in the function, string for jwt tokens and user entity data
func (userService *userService) Login(loginDto dto.LoginDto) (error, string, entity.User) {
	err, user := userService.userRepository.FindByEmail(loginDto.Email)

	if err != nil {
		log.Panic(err.Error())
		return err, "", user
	}

	passwordValid := helper.VerifyPassword(user.Password, loginDto.Password)

	if passwordValid == false {
		return errors.New("bad credentials"), "", user
	}

	err, token := helper.GenerateJWT(user)

	if err != nil {
		log.Panicf("Error on generating jwt (Login) %v", err)
		return err, token, user
	}

	return err, token, user
}

func (userService *userService) Register(registerDto dto.RegisterUserDto) (error, entity.User) {
	var user entity.User

	err, password := helper.HashPassword(registerDto.Password)

	if err != nil {
		return err, user
	}

	user = entity.User{
		Email:       registerDto.Email,
		Password:    password,
		DateOfBirth: registerDto.DateOfBirth.DateString(),
	}

	err, result := userService.userRepository.Create(user)
	if err != nil {
		return err, result
	}
	return nil, result
}

func (userService *userService) GetOne(id string) (error, entity.User) {
	err, user := userService.userRepository.FindById(id)

	if err != nil {
		return err, user
	}

	return nil, user
}
