package repository

import (
	"LoanTrackerApi/internal/entity"
	"context"
)

type UserRepository interface {
	Create(context context.Context, user *entity.User) error
	FindByEmail(context context.Context, email string) (*entity.User, error)
	FindByID(context context.Context, id string) (*entity.GetUserDTO, error)
	Update(context context.Context, user *entity.UpdateUserDTO) error
	GetAllUsers(context context.Context, page, pageSize int) ([]*entity.GetUserDTO, error)
	Verify(context context.Context, id string) error
	ChangePassword(context context.Context, user *entity.User, password string) error
	Delete(context context.Context, id string) error
}
