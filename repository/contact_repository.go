package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web/contact"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(ctx *fiber.Ctx, tx *gorm.DB, contact domain.Contact) domain.Contact
	FindById(ctx *fiber.Ctx, tx *gorm.DB, id int64, userID int) (*domain.Contact, error)
	FindAll(ctx *fiber.Ctx, tx *gorm.DB, userID int, params contact.SearchParams, offset int) ([]domain.Contact, int)
	Update(ctx *fiber.Ctx, tx *gorm.DB, contact *domain.Contact) domain.Contact
	Delete(ctx *fiber.Ctx, tx *gorm.DB, contact *domain.Contact) error
}
