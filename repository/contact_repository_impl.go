package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"gorm.io/gorm"
)

type ContactRepositoryImpl struct {
}

func NewContactRepository() ContactRepository {
	return &ContactRepositoryImpl{}
}

func (repository *ContactRepositoryImpl) Create(ctx *fiber.Ctx, tx *gorm.DB, contact domain.Contact) domain.Contact {
	err := tx.WithContext(ctx.UserContext()).Create(&contact).Error
	helper.PanicIfError(err)
	return contact
}

func (repository *ContactRepositoryImpl) FindById(ctx *fiber.Ctx, tx *gorm.DB, id int64, userID int) (*domain.Contact, error) {
	contact := domain.Contact{}
	err := tx.WithContext(ctx.UserContext()).Where("id = ? AND user_id = ?", id, userID).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (repository *ContactRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *gorm.DB, userID int) []domain.Contact {
	var contacts []domain.Contact
	err := tx.WithContext(ctx.UserContext()).Where("user_id = ?", userID).Find(&contacts).Error
	helper.PanicIfError(err)
	return contacts
}

func (repository *ContactRepositoryImpl) Update(ctx *fiber.Ctx, tx *gorm.DB, contact *domain.Contact) domain.Contact {
	err := tx.WithContext(ctx.UserContext()).Save(contact).Error
	helper.PanicIfError(err)
	return *contact
}

func (repository *ContactRepositoryImpl) Delete(ctx *fiber.Ctx, tx *gorm.DB, contact *domain.Contact) error {
	err := tx.WithContext(ctx.UserContext()).Delete(contact).Error
	return err
}
