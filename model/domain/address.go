package domain

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID         int64 `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	ContactID  int64 `gorm:"column:contact_id"`
	Street     string
	City       string
	Province   string
	Country    string
	PostalCode string
	CreatedAt  time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;autoCreateTime:true;<-:create"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoCreateTime:true;autoUpdateTime:true"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
	Contact    Contact        `gorm:"foreignKey:ContactID;references:ID"`
}
