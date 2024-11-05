package user

import "context"

type UserUseCase interface {
	Add(ctx context.Context, dto AddUserDto) (uint, error)
}

func NewUserUseCase(repo UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

type userUseCase struct {
	repo UserRepository
}

func (u *userUseCase) Add(ctx context.Context, dto AddUserDto) (uint, error) {
	user := dto.MapToModel()

	if err := user.HashPassword(); err != nil {
		return 0, err
	}
	return u.repo.Add(ctx, user)
}
