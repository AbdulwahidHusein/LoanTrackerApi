package account_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AdminDeleteUser(c *gin.Context) {
	userId := c.Param("id")
	err := h.usecase.AdminDeleteUser(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
