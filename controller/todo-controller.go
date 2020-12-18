package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/service"
)

//TodoController is controller for todo
type TodoController interface {
	FindAll(c *gin.Context)
	Create(c *gin.Context)
	FindOne(c *gin.Context)
	Update(c *gin.Context)
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

func (todoController *todoController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todoController.todoService.FindAll(),
	})
}

func (todoController *todoController) Create(c *gin.Context) {
	var createDto dto.CreateTodoDto
	c.ShouldBindJSON(&createDto)
	todoController.todoService.Create(createDto)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
	})
}

func (todoController *todoController) FindOne(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todoController.todoService.FindOne(id),
	})
}

func (todoController *todoController) Update(c *gin.Context) {
	id := c.Param("id")
	var updateDto dto.UpdateTodoDto
	c.ShouldBindJSON(&updateDto)
	todoController.todoService.Update(id, updateDto)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
