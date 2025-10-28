package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/contact"
	"github.com/sorfian/go-contact-management-api/repository"
	"gorm.io/gorm"
)

type ContactServiceImpl struct {
	ContactRepository repository.ContactRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewContactService(contactRepository repository.ContactRepository, DB *gorm.DB, validate *validator.Validate) ContactService {
	return &ContactServiceImpl{ContactRepository: contactRepository, DB: DB, Validate: validate}
}

func (service *ContactServiceImpl) Create(ctx *fiber.Ctx, user domain.User, request *contact.ContactCreateRequest) contact.ContactResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newContact := domain.Contact{
		UserID:    user.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}

	createdContact := service.ContactRepository.Create(ctx, tx, newContact)

	return contact.ContactResponse{
		ID:        createdContact.ID,
		FirstName: createdContact.FirstName,
		LastName:  createdContact.LastName,
		Email:     createdContact.Email,
		Phone:     createdContact.Phone,
	}
}

func (service *ContactServiceImpl) Get(ctx *fiber.Ctx, user domain.User, contactID int64) contact.ContactResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newContact, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	return contact.ContactResponse{
		ID:        newContact.ID,
		FirstName: newContact.FirstName,
		LastName:  newContact.LastName,
		Email:     newContact.Email,
		Phone:     newContact.Phone,
	}
}

func (service *ContactServiceImpl) GetAll(ctx *fiber.Ctx, user domain.User, params contact.SearchParams) contact.SearchResult {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	offset := (params.Page - 1) * params.Size

	// Panggil repository untuk get data dengan filter
	contacts, totalItem := service.ContactRepository.FindAll(ctx, tx, user.ID, params, offset)

	// Hitung total page
	totalPage := (totalItem + params.Size - 1) / params.Size

	var contactResponses []contact.ContactResponse
	for _, newContact := range contacts {
		contactResponses = append(contactResponses, contact.ContactResponse{
			ID:        newContact.ID,
			FirstName: newContact.FirstName,
			LastName:  newContact.LastName,
			Email:     newContact.Email,
			Phone:     newContact.Phone,
		})
	}

	return contact.SearchResult{
		Contacts: contactResponses,
		Paging: web.PagingResponse{
			Page:      params.Page,
			Size:      params.Size,
			TotalPage: totalPage,
			TotalItem: totalItem,
		},
	}
}

func (service *ContactServiceImpl) Update(ctx *fiber.Ctx, user domain.User, contactID int64, request contact.ContactUpdateRequest) contact.ContactResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newContact, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	if request.FirstName != "" {
		newContact.FirstName = request.FirstName
	}

	if request.LastName != "" {
		newContact.LastName = request.LastName
	}

	if request.Email != "" {
		newContact.Email = request.Email
	}

	if request.Phone != "" {
		newContact.Phone = request.Phone
	}

	updatedContact := service.ContactRepository.Update(ctx, tx, newContact)

	return contact.ContactResponse{
		ID:        updatedContact.ID,
		FirstName: updatedContact.FirstName,
		LastName:  updatedContact.LastName,
		Email:     updatedContact.Email,
		Phone:     updatedContact.Phone,
	}
}

func (service *ContactServiceImpl) Delete(ctx *fiber.Ctx, user domain.User, contactID int64) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newContact, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	err = service.ContactRepository.Delete(ctx, tx, newContact)
	helper.PanicIfError(err)
}
