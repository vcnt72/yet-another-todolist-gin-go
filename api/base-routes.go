package api

import (
	"github.com/gin-gonic/gin"
)

// Routes is interface of every routes
type Routes interface {
	RouteGroups(route *gin.Engine)
	noAuthRoutes(route *gin.Engine)
	authRoutes(route *gin.Engine)
}
