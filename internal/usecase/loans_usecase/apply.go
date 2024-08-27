package loans_usecase

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/repository/mongodb/loan_repo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoansUsecase struct {
	loanRepo loan_repo.MongoLoanRepository
}

func NewUsecase(loanRepo loan_repo.MongoLoanRepository) *LoansUsecase {
	return &LoansUsecase{
		loanRepo: loanRepo,
	}
}

func (u *LoansUsecase) ApplyLoan(ctx context.Context, loan *entity.Loan, userId string) error {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		config.Logger.AddLog(ctx, "ApplyLoan", "error occured while converting user with id "+userId+" to object id")
		return err
	}
	loan.IssuerId = objId
	loan.Status = entity.Pending
	config.Logger.AddLog(ctx, "ApplyLoan", "user with Id "+userId+" applied for loan")
	return u.loanRepo.Create(ctx, loan)
}
