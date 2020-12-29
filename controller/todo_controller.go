package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"github.com/yet-another-todo-list-golang/service"
	"log"
	"net/http"
)

//TodoController is controller for todo
type TodoController interface {
	FindAll(c *gin.Context)
	Create(c *gin.Context)
	FindOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type todoController struct {
	todoService service.TodoService
}

//NewTodoController get new controller of todo
func NewTodoController(todoService service.TodoService) TodoController {
	return &todoController{
		todoService: todoService,
	}
}

func (todoController *todoController) FindAll(c *gin.Context) {
	err, todos := todoController.todoService.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todos,
	})
}

func (todoController *todoController) Create(c *gin.Context) {
	var createDto dto.CreateTodoDto
	//log.Print("Vincent")
	err := c.ShouldBindJSON(&createDto)
	if err != nil {
		log.Panic(err.Error())
	}
	val, _ := c.Get("user")
	user := val.(entity.User)
	err = todoController.todoService.Create(createDto, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unknown error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
	})
}

func (todoController *todoController) FindOne(c *gin.Context) {
	id := c.Param("id")
	err, todo := todoController.todoService.FindOne(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todo,
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

func (todoController *todoController) Delete(c *gin.Context) {
	id := c.Param("id")

	todoController.todoService.Delete(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
