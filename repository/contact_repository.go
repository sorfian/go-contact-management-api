package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/model/domain"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(ctx *fiber.Ctx, tx *gorm.DB, contact domain.Contact) domain.Contact
	FindById(ctx *fiber.Ctx, tx *gorm.DB, id int64, userID int) (*domain.Contact, error)
	FindAll(ctx *fiber.Ctx, tx *gorm.DB, userID int) []domain.Contact
	Update(ctx *fiber.Ctx, tx *gorm.DB, contact *domain.Contact) domain.Contact
	Delete(ctx *fiber.Ctx, tx *gorm.DB, contact *domain.Contact) error
}
