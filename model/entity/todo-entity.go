package entity

import "time"

// TodoStatus is enum status of type
type TodoStatus string

// TodoStatus represent if the todo is complete or not
const (
	NotComplete TodoStatus = "NOT_COMPLETE"
	Complete    TodoStatus = "COMPLETE"
)

// Todo is list of things that you wanna do
type Todo struct {
	Base
	Name        string    `json:"name" gorm:"type:varchar(120)" binding:"required,min=1,max=100"`
	Description string    `json:"description" gorm:"type:text" binding:"required"`
	Status      string    `json:"status" gorm:"type:varchar(50);default:NOT_COMPLETE"`
	UserID      string    `json:"-" gorm:"type:uuid""`
	User        User      `json:"user"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}
