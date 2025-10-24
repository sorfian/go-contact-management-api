package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
)

type ContactService interface {
	Create(ctx *fiber.Ctx, user domain.User, request *web.ContactCreateRequest) web.ContactResponse
	Get(ctx *fiber.Ctx, user domain.User, contactID int64) web.ContactResponse
	GetAll(ctx *fiber.Ctx, user domain.User) []web.ContactResponse
	Update(ctx *fiber.Ctx, user domain.User, contactID int64, request web.ContactUpdateRequest) web.ContactResponse
	Delete(ctx *fiber.Ctx, user domain.User, contactID int64)
}
