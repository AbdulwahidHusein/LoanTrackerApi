package loans_usecase

import (
	"LoanTrackerApi/internal/entity"
	"context"
)

type Usecase interface {
	ApplyLoan(ctx context.Context, loan *entity.Loan, userId string) error
	ViewLoan(ctx context.Context, LoanId string, user *entity.User) (*entity.Loan, error)
	ApproveLoan(ctx context.Context, loanId string) (entity.Loan, error)
	RejectLoan(ctx context.Context, loanId string) (entity.Loan, error)
	GetLoans(ctx context.Context, page, pageSize int, filter entity.LoanFilter) ([]*entity.Loan, error)
	DeleteLoan(ctx context.Context, loanId string) error
}
