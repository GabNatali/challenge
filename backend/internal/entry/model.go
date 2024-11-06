package entry

import (
	"time"

	"github.com/codeableorg/weekend-challenge-13-GabNatali/internal/user"
	"gorm.io/gorm"
)

type EntryModel struct {
	Id         uint       `json:"id" gorm:"primaryKey"`
	Title      string     `json:"title" gorm:"not null"`
	Content    string     `json:"content"`
	StatusCode int        `json:"statusCode"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Deleted    gorm.DeletedAt
	UserId     uint           `json:"user_id" gorm:"not null;index"`
	User       user.UserModel `json:"-" gorm:"foreignKey:UserId;references:Id"`
}

func NewEntry(title string, content string, userId uint) EntryModel {
	return EntryModel{
		Title:      title,
		Content:    content,
		StatusCode: 1,
		Status:     "CREATED",
		UserId:     userId,
	}
}
