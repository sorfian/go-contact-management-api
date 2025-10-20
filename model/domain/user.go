package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	Username  string         `gorm:"column:username;unique_index"`
	Password  string         `gorm:"column:password"`
	Name      string         `gorm:"column:name"`
	Token     string         `gorm:"column:token"`
	TokenExp  int64          `gorm:"column:token_exp"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;autoCreateTime:true;<-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoCreateTime:true;autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Contacts  []Contact      `gorm:"foreignKey:UserID;references:ID"`
}

func (user *User) TableName() string {
	return "users"
}
