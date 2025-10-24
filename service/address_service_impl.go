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

type AddressServiceImpl struct {
	AddressRepository repository.AddressRepository
	ContactRepository repository.ContactRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewAddressService(addressRepository repository.AddressRepository, contactRepository repository.ContactRepository, DB *gorm.DB, validate *validator.Validate) AddressService {
	return &AddressServiceImpl{
		AddressRepository: addressRepository,
		ContactRepository: contactRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *AddressServiceImpl) Create(ctx *fiber.Ctx, user domain.User, contactID int64, request *web.AddressCreateRequest) web.AddressResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err = service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	address := domain.Address{
		ContactID:  contactID,
		Street:     request.Street,
		City:       request.City,
		Province:   request.Province,
		Country:    request.Country,
		PostalCode: request.PostalCode,
	}

	createdAddress := service.AddressRepository.Create(ctx, tx, address)

	return web.AddressResponse{
		ID:         createdAddress.ID,
		Street:     createdAddress.Street,
		City:       createdAddress.City,
		Province:   createdAddress.Province,
		Country:    createdAddress.Country,
		PostalCode: createdAddress.PostalCode,
	}
}

func (service *AddressServiceImpl) Get(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64) web.AddressResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	address, err := service.AddressRepository.FindById(ctx, tx, addressID, contactID)
	if err != nil {
		panic(helper.NewNotFoundError("address not found"))
	}

	return web.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		Country:    address.Country,
		PostalCode: address.PostalCode,
	}
}

func (service *AddressServiceImpl) GetAll(ctx *fiber.Ctx, user domain.User, contactID int64) []web.AddressResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	addresses := service.AddressRepository.FindAll(ctx, tx, contactID)

	var addressResponses []web.AddressResponse
	for _, address := range addresses {
		addressResponses = append(addressResponses, web.AddressResponse{
			ID:         address.ID,
			Street:     address.Street,
			City:       address.City,
			Province:   address.Province,
			Country:    address.Country,
			PostalCode: address.PostalCode,
		})
	}

	return addressResponses
}

func (service *AddressServiceImpl) Update(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64, request web.AddressUpdateRequest) web.AddressResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err = service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	address, err := service.AddressRepository.FindById(ctx, tx, addressID, contactID)
	if err != nil {
		panic(helper.NewNotFoundError("address not found"))
	}

	if request.Street != "" {
		address.Street = request.Street
	}

	if request.City != "" {
		address.City = request.City
	}

	if request.Province != "" {
		address.Province = request.Province
	}

	if request.Country != "" {
		address.Country = request.Country
	}

	if request.PostalCode != "" {
		address.PostalCode = request.PostalCode
	}

	updatedAddress := service.AddressRepository.Update(ctx, tx, address)

	return web.AddressResponse{
		ID:         updatedAddress.ID,
		Street:     updatedAddress.Street,
		City:       updatedAddress.City,
		Province:   updatedAddress.Province,
		Country:    updatedAddress.Country,
		PostalCode: updatedAddress.PostalCode,
	}
}

func (service *AddressServiceImpl) Delete(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	address, err := service.AddressRepository.FindById(ctx, tx, addressID, contactID)
	if err != nil {
		panic(helper.NewNotFoundError("address not found"))
	}

	err = service.AddressRepository.Delete(ctx, tx, address)
	helper.PanicIfError(err)
}
