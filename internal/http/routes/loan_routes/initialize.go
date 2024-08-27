package loan_routes

import (
	"LoanTrackerApi/internal/http/handlers/loan_handler"
	"LoanTrackerApi/internal/repository/mongodb/loan_repo"
	"LoanTrackerApi/internal/usecase/loans_usecase"
	"LoanTrackerApi/pkg/mongodb"
)

func GetHandler() *loan_handler.LoanHandler {
	loanCollection := mongodb.GetCollection("loans")

	loanRepo := loan_repo.NewMongoLoanRepository(loanCollection)
	loanUsecase := loans_usecase.NewUsecase(*loanRepo)

	return loan_handler.NewLoanHandler(loanUsecase)
}
