package user

type UserUseCase interface {
	Add(dto AddUserDto) (uint, error)
}

func NewUserUseCase(repo UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

type userUseCase struct {
	repo UserRepository
}

func (u *userUseCase) Add(dto AddUserDto) (uint, error) {
	user, err := dto.MapToModel()
	if err != nil {
		return 0, err
	}

	if err := user.HashPassword(); err != nil {
		return 0, err
	}
	return u.repo.Add(user)
}
