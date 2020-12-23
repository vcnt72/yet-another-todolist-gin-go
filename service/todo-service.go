package service

import (
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/repository"
	"log"
)

// TodoService service of todo
type TodoService interface {
	FindAll() []entity.Todo
	Create(createDto dto.CreateTodoDto, user entity.User) error
	FindOne(id string) entity.Todo
	Update(id string, updateDto dto.UpdateTodoDto)
	Delete(id string)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

// NewTodoService get service of todo
func NewTodoService() TodoService {
	todoRepository := repository.NewTodoRepository()
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (todoService *todoService) FindAll() []entity.Todo {
	return todoService.todoRepository.FindAll()
}

func (todoService *todoService) Create(createDto dto.CreateTodoDto, user entity.User) error {
	todo := &entity.Todo{
		Name:        createDto.Name,
		Description: createDto.Description,
	}

	err := todoService.todoRepository.Create(*todo, user)

	if err != nil {
		log.Panic(err.Error())
		return err
	}
	return nil
}

func (todoService *todoService) FindOne(id string) entity.Todo {
	return todoService.todoRepository.FindOne(id)
}

func (todoService *todoService) Update(id string, updateDto dto.UpdateTodoDto) {
	todo := todoService.FindOne(id)
	todo.Name = updateDto.Name
	todo.Description = updateDto.Description
	todo.Status = updateDto.Status
	todoService.todoRepository.Update(todo)
}

func (todoService *todoService) Delete(id string) {
	todoService.todoRepository.Delete(id)
}
