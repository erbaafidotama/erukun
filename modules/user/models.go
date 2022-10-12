package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UserUuid  uuid.UUID `gorm:"primaryKey;type:uuid" json:"user_uuid"`
	Username  string
	FullName  string
	Password  string
	Email     string
	DateBirth string `json:"date_birth"`
	AdminRole bool   `gorm:"default:0"`
}

//* untuk membuat table pada cutstom schema
// func (UserModel) TableName() string {
// 	return "14.users"
// }

func (UserModel) TableName() string {
	return "master_users" // overide nama table user_models -> users
}
