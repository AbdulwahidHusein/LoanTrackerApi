package loans_usecase

import (
	"LoanTrackerApi/internal/entity"
	"context"
	"errors"
)

func (u *LoansUsecase) ViewLoan(ctx context.Context, LoanId string, user *entity.User) (*entity.Loan, error) {
	loan, err := u.loanRepo.FindByID(ctx, LoanId)
	if err != nil {
		return nil, err
	}
	if loan == nil {
		return nil, errors.New("loan not found")
	}
	if loan.IssuerId != user.ID && user.Role != "admin" {
		return nil, errors.New("unauthorized access")
	}
	return loan, nil
}

// func (u *LoansUsecase) ViewAllLoans(ctx context.Context, user *entity.User) ([]entity.Loan, error) {
