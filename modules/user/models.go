package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UserUuid  uuid.UUID `gorm:"type:uuid" json:"user_uuid"`
	Username  string
	FullName  string
	Password  string
	Email     string
	DateBirth time.Time `json:"date_birth"`
	AdminRole bool      `gorm:"default:0"`
}

func (UserModel) TableName() string {
	return "14.users"
}
