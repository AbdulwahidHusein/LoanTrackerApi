package repository

import "LoanTrackerApi/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	Update(user *entity.User) error
	GetAllUsers() ([]*entity.User, error)
}
