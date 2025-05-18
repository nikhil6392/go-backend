package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhil6392/go-backend/controllers"
)

func Register(router *gin.Engine){
	userGroup := router.Group("/users")
	{
		userGroup.POST("/signup", controllers.SignUp)
	}
}
