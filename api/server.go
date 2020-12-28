package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var todo Routes
var user Routes

func init() {
	todo = NewTodoRoutes()
	user = NewUserRoutes()
}

// Server Setting up server
func Server() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(corsConfig()))
	todo.RouteGroups(server)
	user.RouteGroups(server)
	return server
}

func corsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	return corsConfig
}
