package loan_handler

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/http/handlers"
	"LoanTrackerApi/internal/usecase/loans_usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	useCase loans_usecase.Usecase
}

func NewLoanHandler(useCase loans_usecase.Usecase) *LoanHandler {
	return &LoanHandler{
		useCase: useCase,
	}
}

func (h *LoanHandler) Apply(c *gin.Context) {
	Id, err := handlers.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan := &entity.Loan{}

	if err := c.ShouldBindJSON(loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.useCase.ApplyLoan(c, loan, Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Loan applied successfully"})

}
