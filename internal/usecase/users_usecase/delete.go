package users_usecase

import "context"

func (u Usecase) AdminDeleteUser(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}
