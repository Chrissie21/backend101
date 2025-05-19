package controllers

import (
	"backend101/database"
	"backend101/models"
	"backend101/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateTransaction godoc
// @Summary Create a transaction
// @Description Add a new income or expense transaction
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param transaction body dto.CreateTransactionInput true "Transaction to create"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /transactions [post]
func CreateTransaction(c *gin.Context) {
	var tx models.Transaction

	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate
	if validationErrors := utils.ValidateStruct(&tx); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
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

// GetTransactions godoc
// @Summary Get all user transactions
// @Description Retrieve all transactions for the authenticated user
// @Tags Transactions
// @Produce  json
// @Success 200 {array} models.Transaction
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /transactions [get]
func GetTransactions(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var transaction []models.Transaction
	if err := database.DB.Where("user_id = ?", userID).Find(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction godoc
// @Summary Update a transaction
// @Description Update an existing transaction by ID for the authenticated user
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param id path string true "Transaction ID"
// @Param transaction body dto.UpdateTransactionInput true "Updated transaction data"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /transactions/{id} [put]
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

	// Validate
	if validationErrors := utils.ValidateStruct(&input); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	// Only allow updates to this field and only update fields after validation
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

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete a transaction by ID for the authenticated user
// @Tags Transactions
// @Produce  json
// @Param id path string true "Transaction ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /transactions/{id} [delete]
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

// GetBalance godoc
// @Summary Get current balance
// @Description Calculate and return total income, total expenses, and balance status (positive/negative)
// @Tags Transactions
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /transactions/balance [get]
func GetBalance(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var incomeTotal float64
	var expenseTotal float64

	// Sum incomes
	database.DB.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, "income").
		Select("COALESCE(SUM(amount), 0)").Scan(&incomeTotal)

	// Sum expenses
	database.DB.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, "expense").
		Select("COALESCE(SUM(amount), 0)").Scan(&expenseTotal)

	balance := incomeTotal - expenseTotal
	status := "positive"
	if balance < 0 {
		status = "negative"
	}

	c.JSON(http.StatusOK, gin.H{
		"income_total":   incomeTotal,
		"expense_total":  expenseTotal,
		"balance":        balance,
		"financial_zone": status,
	})
}
