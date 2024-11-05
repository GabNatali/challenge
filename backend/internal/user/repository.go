package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(ctx context.Context, user UserModel) (uint, error)
	GetByEmail(ctx context.Context, email string) (UserModel, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Add(ctx context.Context, user UserModel) (uint, error) {
	tx := r.db.Create(&user)

	if err := tx.Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (UserModel, error) {
	var user UserModel
	tx := r.db.Where("email = ?", email).First(&user)
	return user, tx.Error
}
