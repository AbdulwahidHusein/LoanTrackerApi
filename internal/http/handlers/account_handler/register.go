package account_handler

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/usecase/users_usecase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase users_usecase.UserUsecase
}

func NewHandler(usecase users_usecase.UserUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterationHandler(c *gin.Context) {
	var createUSer entity.CreateUserDTO

	if err := c.ShouldBindJSON(&createUSer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := entity.User{
		Email:    createUSer.Email,
		Password: createUSer.Password,
	}

	if err := h.usecase.Register(context.Background(), &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Succesfully Verifaction Emails sent to your email"})
}
