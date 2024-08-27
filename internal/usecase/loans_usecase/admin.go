package loans_usecase

import (
	"LoanTrackerApi/internal/entity"
	"context"
	"errors"
)

func (u *LoansUsecase) ViewAllLoans(ctx context.Context, page int, pageSize int) ([]*entity.Loan, error) {
	loans, err := u.loanRepo.GetLoans(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (u *LoansUsecase) ApproveLoan(ctx context.Context, loanId string) (entity.Loan, error) {
	loan, err := u.loanRepo.FindByID(ctx, loanId)
	if err != nil {
		return entity.Loan{}, err
	}

	if loan == nil {
		return entity.Loan{}, errors.New("loan not found")
	}
	loan.Status = entity.Approved
	return *loan, u.loanRepo.Update(ctx, loan)
}

func (u *LoansUsecase) RejectLoan(ctx context.Context, loanId string) (entity.Loan, error) {
	loan, err := u.loanRepo.FindByID(ctx, loanId)
	if err != nil {
		return entity.Loan{}, err
	}

	if loan == nil {
		return entity.Loan{}, errors.New("loan not found")
	}
	loan.Status = entity.Rejected
	return *loan, u.loanRepo.Update(ctx, loan)
}
