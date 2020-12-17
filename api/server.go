package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/config"
)

// ServerRun run server
func ServerRun() {
	server := gin.Default()
	todoRoutes(server)
	server.Run(fmt.Sprintf(":%s", config.GetEnvConfig("server.port")))
}
