package loan_handler

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/http/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *LoanHandler) GetLoanHandler(c *gin.Context) {
	loanId := c.Param("id")
	claims, err := handlers.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	objId, err := primitive.ObjectIDFromHex(claims.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := entity.User{
		ID:   objId,
		Role: claims.Role,
	}

	loan, err := h.useCase.ViewLoan(c, loanId, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"loan": loan})

}
