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
	FindOne(id string) entity.Todo
	Update(todo entity.Todo)
	Delete(id string)
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

func (todoRepository *todoRepository) Update(todo entity.Todo) {
	todoRepository.db.Save(&todo)
}

func (todoRepository *todoRepository) FindOne(id string) entity.Todo {
	var todo entity.Todo
	todoRepository.db.First(&todo, "id = ?", id)
	return todo
}

func (todoRepository *todoRepository) Delete(id string) {
	// var todo entity.Todo

	todoRepository.db.Delete(&entity.Todo{}, "id = ?", id)
}
