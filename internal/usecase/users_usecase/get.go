package users_usecase

import (
	"LoanTrackerApi/internal/entity"
	"context"
	"errors"
)

func (u *Usecase) GetMyProfile(ctx context.Context, id string) (entity.GetUserDTO, error) {

	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return entity.GetUserDTO{}, err
	}
	if user == nil {
		return entity.GetUserDTO{}, errors.New("user not found")
	}

	return entity.GetUserDTO{
		ID:       user.ID,
		UserName: user.Profile.FirstName + " " + user.Profile.LastName,
		Email:    user.Email,
		Profile:  user.Profile,
		Role:     user.Role,
		Created:  user.Created,
		Updated:  user.Updated,
		Verified: user.Verified,
	}, nil
}

func (u *Usecase) AdminGetAllUsers(ctx context.Context, page int, pageSize int) ([]*entity.GetUserDTO, error) {

	users, err := u.userRepo.GetAllUsers(ctx, page, pageSize)
	if err != nil {
		return []*entity.GetUserDTO{}, err
	}

	return users, nil
}
