package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/controller"
)

type userRoutes struct {
	userController controller.UserController
}

func NewUserRoutes() Routes {
	return &userRoutes{
		userController: controller.NewUserController(),
	}
}

func (userRoute *userRoutes) RouteGroups(route *gin.Engine) {
	userRoute.noAuthRoutes(route)
	userRoute.authRoutes(route)
}

func (userRoute *userRoutes) noAuthRoutes(route *gin.Engine) {
	routes := route.Group("/api/user")
	{
		routes.POST("/login", func(context *gin.Context) {
			userRoute.userController.Login(context)
		})

		routes.POST("/register", func(context *gin.Context) {
			userRoute.userController.Register(context)
		})
	}
}

func (userRoute *userRoutes) authRoutes(route *gin.Engine) {

}
