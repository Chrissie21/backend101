package routes

import (
	"backend101/controllers"
	"backend101/middleware"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	tx := router.Group("/api/transactions")
	tx.Use(middleware.JWTMiddleware())
	{
		tx.POST("/", controllers.CreateTransaction)
		tx.GET("/", controllers.GetTransactions)
	}
}
