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
	if err := database.DB.Where("user_id = ?", userID).Find(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var tx models.Transaction
	if err := database.DB.Where("id =? AND user_id = ?", id, userID).First(&tx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Only allow updates to this field
	tx.Amount = input.Amount
	tx.Category = input.Category
	tx.Description = input.Description
	tx.Type = input.Type
	tx.Date = input.Date

	if err := database.DB.Save(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func DeleteTransaction(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var tx models.Transaction
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&tx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if err := database.DB.Delete(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}
