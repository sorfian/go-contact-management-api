package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web/contact"
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

func (service *ContactServiceImpl) GetAll(ctx *fiber.Ctx, user domain.User) []contact.ContactResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	contacts := service.ContactRepository.FindAll(ctx, tx, user.ID)

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

	return contactResponses
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
