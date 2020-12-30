package repository

import (
	"github.com/yet-another-todo-list-golang/db"
	"github.com/yet-another-todo-list-golang/model/entity"
	"gorm.io/gorm"
	"log"
)

// TodoRepository is repository for todo db
type TodoRepository interface {
	FindAll() (error, []entity.Todo)
	Create(todo entity.Todo, user entity.User) error
	FindOne(id string) (error, entity.Todo)
	Update(todo entity.Todo) error
	Delete(id string) error
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

func (todoRepository *todoRepository) FindAll() (error, []entity.Todo) {
	var todos []entity.Todo
	err := todoRepository.connection.Joins("User").Find(&todos).Error

	if err != nil {
		log.Fatal(err.Error())
		return err, todos
	}

	return nil, todos
}

func (todoRepository *todoRepository) Create(todo entity.Todo, user entity.User) error {

	err := todoRepository.connection.Model(&user).Association("Todos").Append(&todo)

	if err != nil {
		return err
	}
	return nil
}

func (todoRepository *todoRepository) Update(todo entity.Todo) error {
	err := todoRepository.connection.Save(&todo).Error

	if err != nil {
		log.Panic(err.Error())
		return err
	}

	return nil
}

func (todoRepository *todoRepository) FindOne(id string) (error, entity.Todo) {
	var todo entity.Todo
	err := todoRepository.connection.First(&todo, "id = ?", id).Error
	if err != nil {
		log.Panic(err.Error())
		return err, todo
	}
	return nil, todo
}

func (todoRepository *todoRepository) Delete(id string) error {
	err := todoRepository.connection.Delete(&entity.Todo{}, "id = ?", id).Error
	if err != nil {
		log.Panic(err.Error())
		return err
	}

	return nil
}
