package controllers

import (
	"backend101/database"
	"backend101/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var tx models.Transaction

	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	tx.UserID = userID
	tx.Date = time.Now()

	if err := database.DB.Create(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction."})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func GetTransactions(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var transaction []models.Transaction
	if err := database.DB.Where("user_is = ?", userID).Find(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
