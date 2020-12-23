package repository

import (
	"github.com/yet-another-todo-list-golang/config"
	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/gorm"
	"log"
)

// TodoRepository is repository for todo db
type TodoRepository interface {
	FindAll() []entity.Todo
	Create(todo entity.Todo, user entity.User) error
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
	todoRepository.db.Joins("User").Find(&todos)
	return todos
}

func (todoRepository *todoRepository) Create(todo entity.Todo, user entity.User) error {

	err := todoRepository.db.Model(&user).Association("Todos").Append(&todo)

	if err != nil {
		log.Panic(err.Error())
		return err
	}
	return nil
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
