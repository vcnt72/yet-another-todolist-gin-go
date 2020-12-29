package dto

import "github.com/rickb777/date"

type RegisterUserDto struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth date.Date `json:"dob"`
}
