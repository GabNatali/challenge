package auth

import "github.com/codeableorg/weekend-challenge-13-GabNatali/internal/user"

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggedUserDto struct {
	user.UserDto
	Token string `json:"token"`
}

func (dto LoggedUserDto) MapFromModel(model user.UserModel, token string) LoggedUserDto {
	dto.Id = model.Id
	dto.FirstName = model.FirstName
	dto.LastName = model.LastName
	dto.Email = model.Email
	dto.Token = token

	return dto
}
