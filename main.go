package main

import (
	"backend101/config"
	"backend101/database"
	"backend101/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectPostgres()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.TransactionRoutes(r)

	port := config.Get("PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
