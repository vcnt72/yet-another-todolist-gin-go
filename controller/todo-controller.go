package controller

import (
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/service"
)

type TodoController interface {
	FindAll() []entity.Todo
	Create(createDto *dto.CreateTodoDto)
}

type todoController struct {
	todoService service.TodoService
}

//NewTodoController get new controller of todo
func NewTodoController() TodoController {
	todoService := service.NewTodoService()

	return &todoController{
		todoService: todoService,
	}
}

func (todoController *todoController) FindAll() []entity.Todo {
	return todoController.todoService.FindAll()
}

func (todoController *todoController) Create(createDto *dto.CreateTodoDto) {
	todoController.todoService.Create(*createDto)
}
