package repository

import (
	"github.com/yet-another-todo-list-golang/db"
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
	connection *gorm.DB
}

//NewTodoRepository get new todo repository
func NewTodoRepository() TodoRepository {
	connection := db.GetConnection()
	return &todoRepository{
		connection: connection,
	}
}

func (todoRepository *todoRepository) FindAll() []entity.Todo {
	var todos []entity.Todo
	todoRepository.connection.Joins("User").Find(&todos)
	return todos
}

func (todoRepository *todoRepository) Create(todo entity.Todo, user entity.User) error {

	err := todoRepository.connection.Model(&user).Association("Todos").Append(&todo)

	if err != nil {
		log.Panic(err.Error())
		return err
	}
	return nil
}

func (todoRepository *todoRepository) Update(todo entity.Todo) {
	todoRepository.connection.Save(&todo)
}

func (todoRepository *todoRepository) FindOne(id string) entity.Todo {
	var todo entity.Todo
	todoRepository.connection.First(&todo, "id = ?", id)
	return todo
}

func (todoRepository *todoRepository) Delete(id string) {
	todoRepository.connection.Delete(&entity.Todo{}, "id = ?", id)
}
