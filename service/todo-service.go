package service

import (
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/repository"
)

// TodoService service of todo
type TodoService interface {
	FindAll() []entity.Todo
	Create(createDto dto.CreateTodoDto)
	FindOne(id string) entity.Todo
	Update(id string, updateDto dto.UpdateTodoDto)
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

func (todoService *todoService) Create(createDto dto.CreateTodoDto) {
	todo := &entity.Todo{
		Name:        createDto.Name,
		Description: createDto.Description,
	}

	todoService.todoRepository.Create(*todo)
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
