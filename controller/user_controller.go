package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/service"
	"log"
	"net/http"
)

type UserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (userController *userController) Register(c *gin.Context) {
	var registerDto dto.RegisterUserDto

	err := c.ShouldBindJSON(&registerDto)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "Unknown error",
		})
		return
	}
	err, user := userController.userService.Register(registerDto)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Unknown error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": gin.H{
			"user": user,
		},
	})
	return
}

func (userController *userController) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Unknown error",
		})
		return
	}

	err, token, user := userController.userService.Login(loginDto)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
	return
}
