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

// Server Setting up server
func Server() *gin.Engine {
	config.Init()
	server := gin.Default()
	server.Use(cors.Default())
	todo.RouteGroups(server)
	user.RouteGroups(server)
	return server
}
