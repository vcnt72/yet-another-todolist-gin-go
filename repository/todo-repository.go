package repository

import (
	"github.com/yet-another-todo-list-golang/db"
	"github.com/yet-another-todo-list-golang/model/entity"
)

// TodoRepository is repository for todo db
type TodoRepository interface {
	FindAll() []entity.Todo
	Create(todo entity.Todo)
}

type todoRepository struct {
	db *db.Database
}

//NewTodoRepository get new todo repository
func NewTodoRepository() TodoRepository {
	db := db.Connect()
	return &todoRepository{
		db: db,
	}
}

func (todoRepository *todoRepository) FindAll() []entity.Todo {
	var todos []entity.Todo
	todoRepository.db.Connection.Find(&todos)
	return todos
}

func (todoRepository *todoRepository) Create(todo entity.Todo) {
	todoRepository.db.Connection.Create(&todo)
}
