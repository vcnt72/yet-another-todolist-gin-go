package service

import (
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/repository"
	"log"
)

// TodoService service of todo
type TodoService interface {
	FindAll() (error, []entity.Todo)
	Create(createDto dto.CreateTodoDto, user entity.User) error
	FindOne(id string) (error, entity.Todo)
	Update(id string, updateDto dto.UpdateTodoDto) error
	Delete(id string) error
}

type todoService struct {
	todoRepository repository.TodoRepository
}

// NewTodoService get service of Todo
func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (todoService *todoService) FindAll() (error, []entity.Todo) {
	err, todos := todoService.todoRepository.FindAll()

	if err != nil {
		log.Fatal(err.Error())
		return err, todos
	}

	return nil, todos
}

func (todoService *todoService) Create(createDto dto.CreateTodoDto, user entity.User) error {
	todo := entity.Todo{
		Name:        createDto.Name,
		Description: createDto.Description,
	}

	err := todoService.todoRepository.Create(todo, user)

	if err != nil {
		return err
	}

	return nil
}

func (todoService *todoService) FindOne(id string) (error, entity.Todo) {
	err, todo := todoService.todoRepository.FindOne(id)

	if err != nil {
		return err, todo
	}

	return nil, todo
}

func (todoService *todoService) Update(id string, updateDto dto.UpdateTodoDto) error {
	err, todo := todoService.FindOne(id)
	if err != nil {
		return err
	}
	todo.Name = updateDto.Name
	todo.Description = updateDto.Description
	todo.Status = updateDto.Status
	err = todoService.todoRepository.Update(todo)

	if err != nil {
		return err
	}
	return nil
}

func (todoService *todoService) Delete(id string) error {
	err := todoService.todoRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
