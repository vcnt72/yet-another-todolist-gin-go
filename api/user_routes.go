package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/controller"
	"github.com/yet-another-todo-list-golang/repository"
	"github.com/yet-another-todo-list-golang/service"
)

type userRoutes struct {
	userController controller.UserController
}

func NewUserRoutes() Routes {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return &userRoutes{
		userController: userController,
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
