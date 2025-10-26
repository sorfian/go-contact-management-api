package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(ctx *fiber.Ctx, tx *gorm.DB, address domain.Address) domain.Address
	FindById(ctx *fiber.Ctx, tx *gorm.DB, id int64, contactID int64) (*domain.Address, error)
	FindAll(ctx *fiber.Ctx, tx *gorm.DB, contactID int64) []domain.Address
	Update(ctx *fiber.Ctx, tx *gorm.DB, address *domain.Address) domain.Address
	Delete(ctx *fiber.Ctx, tx *gorm.DB, address *domain.Address) error
}
