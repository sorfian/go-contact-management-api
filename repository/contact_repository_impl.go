package repository

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web/contact"
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
	contactEntity := domain.Contact{}
	err := tx.WithContext(ctx.UserContext()).Where("id = ? AND user_id = ?", id, userID).First(&contactEntity).Error
	if err != nil {
		return nil, err
	}
	return &contactEntity, nil
}

func (repository *ContactRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *gorm.DB, userID int, params contact.SearchParams, offset int) ([]domain.Contact, int) {
	var contacts []domain.Contact
	var totalItem int64

	// Base query dengan user filter
	query := tx.WithContext(ctx.UserContext()).Where("user_id = ?", userID)

	// Tambahkan filter search jika ada
	if params.Name != "" {
		// Search di first_name dan last_name
		query = query.Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ?",
			"%"+strings.ToLower(params.Name)+"%",
			"%"+strings.ToLower(params.Name)+"%")
	}

	if params.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+params.Phone+"%")
	}

	if params.Email != "" {
		query = query.Where("LOWER(email) LIKE ?", "%"+strings.ToLower(params.Email)+"%")
	}

	// Hitung total item sebelum pagination
	query.Model(&domain.Contact{}).Count(&totalItem)

	// Apply pagination dan get data
	err := query.Offset(offset).Limit(params.Size).Find(&contacts).Error
	helper.PanicIfError(err)
	return contacts, int(totalItem)
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
