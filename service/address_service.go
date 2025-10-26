package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web/address"
)

type AddressService interface {
	Create(ctx *fiber.Ctx, user domain.User, contactID int64, request *address.AddressCreateRequest) address.AddressResponse
	Get(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64) address.AddressResponse
	GetAll(ctx *fiber.Ctx, user domain.User, contactID int64) []address.AddressResponse
	Update(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64, request address.AddressUpdateRequest) address.AddressResponse
	Delete(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64)
}
