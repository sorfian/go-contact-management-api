package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/address"
	"github.com/sorfian/go-contact-management-api/service"
)

type AddressControllerImpl struct {
	AddressService service.AddressService
}

func NewAddressController(addressService service.AddressService) AddressController {
	return &AddressControllerImpl{AddressService: addressService}
}

func (controller *AddressControllerImpl) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	request := address.AddressCreateRequest{}
	err = ctx.BodyParser(&request)
	helper.PanicIfError(err)

	addressResponse := controller.AddressService.Create(ctx, *user, contactID, &request)

	webResponse := web.Response{
		Code:   201,
		Status: "Created",
		Data:   addressResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *AddressControllerImpl) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	addressID, err := strconv.ParseInt(ctx.Params("addressId"), 10, 64)
	helper.PanicIfError(err)

	addressResponse := controller.AddressService.Get(ctx, *user, contactID, addressID)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *AddressControllerImpl) GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	addressResponses := controller.AddressService.GetAll(ctx, *user, contactID)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   addressResponses,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *AddressControllerImpl) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	addressID, err := strconv.ParseInt(ctx.Params("addressId"), 10, 64)
	helper.PanicIfError(err)

	var request address.AddressUpdateRequest
	err = ctx.BodyParser(&request)
	helper.PanicIfError(err)

	addressResponse := controller.AddressService.Update(ctx, *user, contactID, addressID, request)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   addressResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *AddressControllerImpl) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*domain.User)

	contactID, err := strconv.ParseInt(ctx.Params("contactId"), 10, 64)
	helper.PanicIfError(err)

	addressID, err := strconv.ParseInt(ctx.Params("addressId"), 10, 64)
	helper.PanicIfError(err)

	controller.AddressService.Delete(ctx, *user, contactID, addressID)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   "Address deleted successfully",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
