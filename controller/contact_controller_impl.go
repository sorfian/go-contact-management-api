package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
	"github.com/sorfian/go-todo-list/service"
)

type ContactControllerImpl struct {
	ContactService service.ContactService
}

func NewContactController(contactService service.ContactService) ContactController {
	return &ContactControllerImpl{ContactService: contactService}
}

func (controller *ContactControllerImpl) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	request := web.ContactCreateRequest{}
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	contactResponse := controller.ContactService.Create(ctx, *user, &request)

	webResponse := web.Response{
		Code:   201,
		Status: "Created",
		Data:   contactResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ContactControllerImpl) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	contactResponse := controller.ContactService.Get(ctx, *user, contactID)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   contactResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ContactControllerImpl) GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactResponses := controller.ContactService.GetAll(ctx, *user)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   contactResponses,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ContactControllerImpl) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	var request web.ContactUpdateRequest
	err = ctx.BodyParser(&request)
	helper.PanicIfError(err)

	contactResponse := controller.ContactService.Update(ctx, *user, contactID, request)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   contactResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ContactControllerImpl) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	controller.ContactService.Delete(ctx, *user, contactID)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   "Contact deleted successfully",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
