package account_handler

import (
	"LoanTrackerApi/internal/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RequestVerifyEmail(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user := entity.LoginUserDTO{
		Email: request.Email,
	}

	go func() {
		err := h.usecase.RequestEmailVerification(user)
		if err != nil {
			log.Printf("Error sending verification email: %v", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Verification email is being sent"})
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing email or token"})
		return
	}

	err := h.usecase.VerifyEmail(token)
	if err != nil {
		log.Printf("Error verifying email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
