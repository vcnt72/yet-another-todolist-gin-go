package repository

import (
	"github.com/yet-another-todo-list-golang/config"
	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/gorm"
)

// TodoRepository is repository for todo db
type TodoRepository interface {
	FindAll() []entity.Todo
	Create(todo entity.Todo)
}

type todoRepository struct {
	db *gorm.DB
}

//NewTodoRepository get new todo repository
func NewTodoRepository() TodoRepository {
	db := config.DatabaseConnect()
	return &todoRepository{
		db: db,
	}
}

func (todoRepository *todoRepository) FindAll() []entity.Todo {
	var todos []entity.Todo
	todoRepository.db.Find(&todos)
	return todos
}

func (todoRepository *todoRepository) Create(todo entity.Todo) {
	todoRepository.db.Create(&todo)
}
