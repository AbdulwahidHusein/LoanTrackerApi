package account_handler

import (
	"LoanTrackerApi/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LoginHandler(c *gin.Context) {

	user := entity.LoginUserDTO{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.usecase.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
