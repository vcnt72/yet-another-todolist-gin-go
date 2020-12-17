package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/controller"
)

var todoController controller.TodoController

func init() {
	todoController = controller.NewTodoController()
}

func todoRoutes(route *gin.Engine) {
	routes := route.Group("/api")
	{
		routes.GET("/todo", func(c *gin.Context) {
			todoController.FindAll(c)
		})

		routes.POST("/todo", func(c *gin.Context) {
			todoController.Create(c)
		})
	}
}
