package entity

import (
	"github.com/rickb777/date"
)

//User is entity of users in database
type User struct {
	Base
	Email       string          `json:"email" gorm:"type:varchar(255)"`
	Password    string          `json:"-" gorm:"type:varchar(255)"`
	DateOfBirth date.DateString `json:"dob" gorm:"type:timestamptz"`
	Todos       []Todo          `json:"-"  json:"todos"`
}
