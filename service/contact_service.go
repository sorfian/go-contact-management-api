package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web/contact"
)

type ContactService interface {
	Create(ctx *fiber.Ctx, user domain.User, request *contact.ContactCreateRequest) contact.ContactResponse
	Get(ctx *fiber.Ctx, user domain.User, contactID int64) contact.ContactResponse
	GetAll(ctx *fiber.Ctx, user domain.User, param contact.SearchParams) contact.SearchResult
	Update(ctx *fiber.Ctx, user domain.User, contactID int64, request contact.ContactUpdateRequest) contact.ContactResponse
	Delete(ctx *fiber.Ctx, user domain.User, contactID int64)
}
