package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"first_name" gorm:"type:char(50);not null"`
	LastName  string         `json:"last_name" gorm:"type:char(50);not null"`
	Email     string         `json:"email" gorm:"type:char(50);not null;unique"`
	Password  string         `json:"password" gorm:"type:char(150);not null"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	Deleted   gorm.DeletedAt `json:"-"`
}

func NewUser(firstName, lastName, email, password string) UserModel {
	return UserModel{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

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
