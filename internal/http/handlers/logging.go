package handlers

import (
	"LoanTrackerApi/config"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {
	// category := c.Query("category")
	logs, err := config.Logger.GetLogs(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
