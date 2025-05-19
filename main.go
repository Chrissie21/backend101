package main

import (
	"backend101/config"
	"backend101/database"
	"backend101/docs"
	"backend101/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadConfig()
	database.ConnectPostgres()

	r := gin.Default()

	//Swagger info
	docs.SwaggerInfo.Title = "Expense Tracker APIs"
	docs.SwaggerInfo.Description = "API documentation for the Expense Tracker backend in Go"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.TransactionRoutes(r)

	// Swagger Docs Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	port := config.Get("PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
