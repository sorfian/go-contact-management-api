package domain

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	UserID    string `gorm:"column:user_id"`
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;autoCreateTime:true;<-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoCreateTime:true;autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	User      User           `gorm:"foreignKey:UserID;references:ID"`
	Addresses []Address      `gorm:"foreignKey:ContactID;references:ID"`
}
