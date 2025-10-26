package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepository() AddressRepository {
	return &AddressRepositoryImpl{}
}

func (repository *AddressRepositoryImpl) Create(ctx *fiber.Ctx, tx *gorm.DB, address domain.Address) domain.Address {
	err := tx.WithContext(ctx.UserContext()).Create(&address).Error
	helper.PanicIfError(err)
	return address
}

func (repository *AddressRepositoryImpl) FindById(ctx *fiber.Ctx, tx *gorm.DB, id int64, contactID int64) (*domain.Address, error) {
	address := domain.Address{}
	err := tx.WithContext(ctx.UserContext()).Where("id = ? AND contact_id = ?", id, contactID).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (repository *AddressRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *gorm.DB, contactID int64) []domain.Address {
	var addresses []domain.Address
	err := tx.WithContext(ctx.UserContext()).Where("contact_id = ?", contactID).Find(&addresses).Error
	helper.PanicIfError(err)
	return addresses
}

func (repository *AddressRepositoryImpl) Update(ctx *fiber.Ctx, tx *gorm.DB, address *domain.Address) domain.Address {
	err := tx.WithContext(ctx.UserContext()).Save(address).Error
	helper.PanicIfError(err)
	return *address
}

func (repository *AddressRepositoryImpl) Delete(ctx *fiber.Ctx, tx *gorm.DB, address *domain.Address) error {
	err := tx.WithContext(ctx.UserContext()).Delete(address).Error
	return err
}
