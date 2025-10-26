package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web/address"
	"github.com/sorfian/go-contact-management-api/repository"
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

func (service *AddressServiceImpl) Create(ctx *fiber.Ctx, user domain.User, contactID int64, request *address.AddressCreateRequest) address.AddressResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err = service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	newAddress := domain.Address{
		ContactID:  contactID,
		Street:     request.Street,
		City:       request.City,
		Province:   request.Province,
		Country:    request.Country,
		PostalCode: request.PostalCode,
	}

	createdAddress := service.AddressRepository.Create(ctx, tx, newAddress)

	return address.AddressResponse{
		ID:         createdAddress.ID,
		Street:     createdAddress.Street,
		City:       createdAddress.City,
		Province:   createdAddress.Province,
		Country:    createdAddress.Country,
		PostalCode: createdAddress.PostalCode,
	}
}

func (service *AddressServiceImpl) Get(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64) address.AddressResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to user
	_, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	addressEntity, err := service.AddressRepository.FindById(ctx, tx, addressID, contactID)
	if err != nil {
		panic(helper.NewNotFoundError("address not found"))
	}

	return address.AddressResponse{
		ID:         addressEntity.ID,
		Street:     addressEntity.Street,
		City:       addressEntity.City,
		Province:   addressEntity.Province,
		Country:    addressEntity.Country,
		PostalCode: addressEntity.PostalCode,
	}
}

func (service *AddressServiceImpl) GetAll(ctx *fiber.Ctx, user domain.User, contactID int64) []address.AddressResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to a user
	_, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	addresses := service.AddressRepository.FindAll(ctx, tx, contactID)

	var addressResponses []address.AddressResponse
	for _, newAddress := range addresses {
		addressResponses = append(addressResponses, address.AddressResponse{
			ID:         newAddress.ID,
			Street:     newAddress.Street,
			City:       newAddress.City,
			Province:   newAddress.Province,
			Country:    newAddress.Country,
			PostalCode: newAddress.PostalCode,
		})
	}

	return addressResponses
}

func (service *AddressServiceImpl) Update(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64, request address.AddressUpdateRequest) address.AddressResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Verify contact belongs to a user
	_, err = service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	addressEntity, err := service.AddressRepository.FindById(ctx, tx, addressID, contactID)
	if err != nil {
		panic(helper.NewNotFoundError("address not found"))
	}

	if request.Street != "" {
		addressEntity.Street = request.Street
	}

	if request.City != "" {
		addressEntity.City = request.City
	}

	if request.Province != "" {
		addressEntity.Province = request.Province
	}

	if request.Country != "" {
		addressEntity.Country = request.Country
	}

	if request.PostalCode != "" {
		addressEntity.PostalCode = request.PostalCode
	}

	updatedAddress := service.AddressRepository.Update(ctx, tx, addressEntity)

	return address.AddressResponse{
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

	// Verify contact belongs to a user
	_, err := service.ContactRepository.FindById(ctx, tx, contactID, user.ID)
	if err != nil {
		panic(helper.NewNotFoundError("contact not found"))
	}

	addressEntity, err := service.AddressRepository.FindById(ctx, tx, addressID, contactID)
	if err != nil {
		panic(helper.NewNotFoundError("address not found"))
	}

	err = service.AddressRepository.Delete(ctx, tx, addressEntity)
	helper.PanicIfError(err)
}
