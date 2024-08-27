package account_handler

import (
	"LoanTrackerApi/pkg/jwt_utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context) {
	access := c.GetHeader("Authorization")
	if access == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access token is required"})
		return
	}

	token, err := jwt_utils.RefreshToken(access)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)

}
