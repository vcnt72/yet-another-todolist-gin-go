package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/controller"
)

type todoRoutes struct {
	todoController controller.TodoController
}

// NewTodoRoutes get new instance of todo routes
func NewTodoRoutes() Routes {
	return &todoRoutes{
		todoController: controller.NewTodoController(),
	}
}

func (tdr *todoRoutes) RouteGroups(route *gin.Engine) {
	tdr.noAuthRoutes(route)
	tdr.authRoutes(route)
}

func (tdr *todoRoutes) noAuthRoutes(route *gin.Engine) {
	routes := route.Group("/api")
	{
		routes.GET("/todo", func(c *gin.Context) {
			tdr.todoController.FindAll(c)
		})

		routes.POST("/todo", func(c *gin.Context) {
			tdr.todoController.Create(c)
		})
	}
}

func (tdr *todoRoutes) authRoutes(route *gin.Engine) {

}
