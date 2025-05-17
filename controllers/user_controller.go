package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	c.JSON(http.StatusOK, gin.H{
		"message": "You are authenticated",
		"user_id": userID,
	})
}
