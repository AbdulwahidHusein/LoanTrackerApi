package account_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ResetPasswordRequest(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.usecase.RequestPasswordResetUsecase(request.Email)
	if err != nil {
		log.Printf("Error sending password reset email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send password reset email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	var request struct {
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = h.usecase.ResetPassword(token, request.Password)
	if err != nil {
		log.Printf("Error resetting password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})

}
