package loans_usecase

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"
	"context"
	"errors"
)

func (u *LoansUsecase) GetLoans(ctx context.Context, page int, pageSize int, filter entity.LoanFilter) ([]*entity.Loan, error) {
	loans, err := u.loanRepo.GetLoans(ctx, page, pageSize, filter)
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
	config.Logger.AddLog(ctx, "ApproveLoan", "loan with id "+loanId+" approved")
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
	config.Logger.AddLog(ctx, "RejectLoan", "loan with id "+loanId+" rejected")
	return *loan, u.loanRepo.Update(ctx, loan)
}

func (u *LoansUsecase) DeleteLoan(ctx context.Context, loanId string) error {

	err := u.loanRepo.Delete(ctx, loanId)
	if err != nil {
		config.Logger.AddLog(ctx, "DeleteLoan", "error occured while deleting loan with id "+loanId)
		return err
	}
	config.Logger.AddLog(ctx, "DeleteLoan", "loan with id "+loanId+" deleted")
	return nil
}
