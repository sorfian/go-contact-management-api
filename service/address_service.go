package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
)

type AddressService interface {
	Create(ctx *fiber.Ctx, user domain.User, contactID int64, request *web.AddressCreateRequest) web.AddressResponse
	Get(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64) web.AddressResponse
	GetAll(ctx *fiber.Ctx, user domain.User, contactID int64) []web.AddressResponse
	Update(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64, request web.AddressUpdateRequest) web.AddressResponse
	Delete(ctx *fiber.Ctx, user domain.User, contactID int64, addressID int64)
}
