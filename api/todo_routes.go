package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/api/middleware"
	"github.com/yet-another-todo-list-golang/controller"
	"github.com/yet-another-todo-list-golang/repository"
	"github.com/yet-another-todo-list-golang/service"
)

type todoRoutes struct {
	todoController controller.TodoController
}

// NewTodoRoutes get new instance of todo routes
func NewTodoRoutes() Routes {
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)
	return &todoRoutes{
		todoController: todoController,
	}
}

func (tdr *todoRoutes) RouteGroups(route *gin.Engine) {
	tdr.noAuthRoutes(route)
	tdr.authRoutes(route)
}

func (tdr *todoRoutes) noAuthRoutes(route *gin.Engine) {

}

func (tdr *todoRoutes) authRoutes(route *gin.Engine) {
	routes := route.Group("/api", middleware.JwtAuth())
	{
		routes.GET("/todo", func(c *gin.Context) {
			tdr.todoController.FindAll(c)
		})

		routes.POST("/todo", func(c *gin.Context) {
			tdr.todoController.Create(c)
		})

		routes.GET("/todo/:id", func(c *gin.Context) {
			tdr.todoController.FindOne(c)
		})

		routes.PUT("/todo/:id", func(c *gin.Context) {
			tdr.todoController.Update(c)
		})

		routes.DELETE("/todo/:id", func(c *gin.Context) {
			tdr.todoController.Delete(c)
		})
	}
}
