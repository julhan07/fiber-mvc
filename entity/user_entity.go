package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UserEntity) TableName() string {
	return "users"
}

func (u *UserEntity) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
