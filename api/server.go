package api

import (
	"github.com/gin-gonic/gin"
)

// ServerRun run server
func ServerRun() {
	server := gin.Default()
	todoRoutes(server)
	server.Run(":8080")
}
