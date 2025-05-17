package routes

import (
	"backend101/controllers"
	"backend101/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("api/user")
	user.Use(middleware.JWTMiddleware())
	{
		user.GET("/me", controllers.Me)
	}
}
