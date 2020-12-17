package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/config"
)

var todo Routes

func init() {
	todo = NewTodoRoutes()
}

// ServerRun run server
func ServerRun() {
	server := gin.Default()
	todo.RouteGroups(server)
	server.Run(":" + config.GetEnvConfig("server.port"))
}
