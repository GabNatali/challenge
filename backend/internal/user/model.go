package user

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	Id        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"type:char(50);not null"`
	LastName  string `gorm:"type:char(50);not null"`
	Email     string `gorm:"type:char(50);not null;unique"`
	Password  string `gorm:"type:char(150);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Deleted   gorm.DeletedAt
}

func NewUser(firstName, lastName, email, password string) (UserModel, error) {
	user := UserModel{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	if err := user.Validate(); err != nil {
		return UserModel{}, err
	}

	return user, nil

}

func (user *UserModel) HashPassword() error {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)

	return nil
}

func (user *UserModel) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (user *UserModel) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.FirstName, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.LastName, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 8)),
	)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
