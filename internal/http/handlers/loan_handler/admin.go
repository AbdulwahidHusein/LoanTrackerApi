package loan_handler

import (
	"LoanTrackerApi/internal/entity"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *LoanHandler) GetLoans(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Parse filter parameters
	filter := entity.LoanFilter{}
	if status := c.Query("status"); status != "" {
		filter.Status = status
	}

	// Parse sorting parameters
	orderBy := c.DefaultQuery("orderBy", "date_asc") // Default sorting order
	filter.OrderBy = []string{orderBy}

	loans, err := h.useCase.GetLoans(c.Request.Context(), page, pageSize, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loans": loans})
}

func (h *LoanHandler) ApproveLoanHanlder(c *gin.Context) {
	loanId := c.Param("id")
	loan, err := h.useCase.ApproveLoan(context.Background(), loanId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Loan approved successfully", "loan": loan})
}

func (h *LoanHandler) RejectLoanHandler(c *gin.Context) {
	loanId := c.Param("id")
	loan, err := h.useCase.RejectLoan(context.Background(), loanId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan rejected successfully", "loan": loan})
}

func (h *LoanHandler) DeleteLoanHandler(c *gin.Context) {
	loanId := c.Param("id")
	err := h.useCase.DeleteLoan(context.Background(), loanId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}
