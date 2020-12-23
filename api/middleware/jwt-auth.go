package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yet-another-todo-list-golang/helper"
	"github.com/yet-another-todo-list-golang/service"
	"gorm.io/gorm"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader("Authorization")

		if len(authorization) == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		token := authorization[len("Bearer "):]

		err, jwtObj := helper.VerifyJWT(token)
		claims := jwtObj.Claims.(*jwt.StandardClaims)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})

				return
			}
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unknown error",
			})
			return
		}

		if !jwtObj.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		err, user := service.NewUserService().GetOne(claims.Id)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				return
			}
		}
		context.Set("user", user)
	}
}
