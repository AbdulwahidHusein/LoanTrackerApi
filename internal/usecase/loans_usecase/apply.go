package loans_usecase

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoansUsecase struct {
	loanRepo repository.LoanRepository
}

func NewUsecase(loanRepo repository.LoanRepository) *LoansUsecase {
	return &LoansUsecase{
		loanRepo: loanRepo,
	}
}

func (u *LoansUsecase) ApplyLoan(ctx context.Context, loan *entity.Loan, userId string) error {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	loan.IssuerId = objId
	loan.Status = entity.Pending
	return u.loanRepo.Create(ctx, loan)
}
