package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
	"github.com/sorfian/go-todo-list/repository"
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

func (service *ContactServiceImpl) Create(ctx *fiber.Ctx, user domain.User, request *web.ContactCreateRequest) web.ContactResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	contact := domain.Contact{
		UserID:    user.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}

	createdContact := service.ContactRepository.Create(ctx, tx, contact)

	return web.ContactResponse{
		ID:        createdContact.ID,
		FirstName: createdContact.FirstName,
		LastName:  createdContact.LastName,
		Email:     createdContact.Email,
		Phone:     createdContact.Phone,
	}
}

func (service *ContactServiceImpl) Get(ctx *fiber.Ctx, user domain.User, contactID int64) web.ContactResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	contact, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	return web.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}
}

func (service *ContactServiceImpl) GetAll(ctx *fiber.Ctx, user domain.User) []web.ContactResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	contacts := service.ContactRepository.FindAll(ctx, tx, user.ID)

	var contactResponses []web.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, web.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email,
			Phone:     contact.Phone,
		})
	}

	return contactResponses
}

func (service *ContactServiceImpl) Update(ctx *fiber.Ctx, user domain.User, contactID int64, request web.ContactUpdateRequest) web.ContactResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	contact, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	if request.FirstName != "" {
		contact.FirstName = request.FirstName
	}

	if request.LastName != "" {
		contact.LastName = request.LastName
	}

	if request.Email != "" {
		contact.Email = request.Email
	}

	if request.Phone != "" {
		contact.Phone = request.Phone
	}

	updatedContact := service.ContactRepository.Update(ctx, tx, contact)

	return web.ContactResponse{
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

	contact, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	err = service.ContactRepository.Delete(ctx, tx, contact)
	helper.PanicIfError(err)
}
