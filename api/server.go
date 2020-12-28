package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/config"
)

var todo Routes
var user Routes

func init() {
	todo = NewTodoRoutes()
	user = NewUserRoutes()
}

// ServerRun run server
func ServerRun() {
	server := gin.Default()
	server.Use(cors.Default())
	todo.RouteGroups(server)
	user.RouteGroups(server)
	server.Run(":" + config.GetEnvConfig("server.port"))
}
