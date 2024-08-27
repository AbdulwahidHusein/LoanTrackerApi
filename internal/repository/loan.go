package repository

import (
	"LoanTrackerApi/internal/entity"
	"context"
)

type LoanRepository interface {
	Create(context context.Context, loan *entity.Loan) error
	FindByID(context context.Context, id string) (*entity.Loan, error)
	FindByUserID(context context.Context, id string) ([]*entity.Loan, error)
	Update(context context.Context, loan *entity.Loan) error
	Delete(context context.Context, id string) error
	GetLoans(context context.Context, page, pageSize int, filter entity.LoanFilter) ([]*entity.Loan, error)
}
