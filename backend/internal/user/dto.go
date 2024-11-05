package user

type UserDto struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (dto UserDto) MapFromModel(user UserModel) UserDto {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email

	return dto
}

type AddUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (dto AddUserDto) MapToModel() UserModel {
	return NewUser(
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
	)
}
