package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/config"
)

// ServerRun run server
func ServerRun() {
	server := gin.Default()
	todoRoutes(server)
	server.Run(":" + config.GetEnvConfig("server.port"))
}
