package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/user"
	"github.com/sorfian/go-contact-management-api/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
func (controller *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	request := user.UserRegisterRequest{}
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	tokenResponse := controller.UserService.Register(ctx, &request)

	webResponse := web.Response{
		Code:   201,
		Status: "Created",
		Data:   tokenResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	request := user.UserLoginRequest{}
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	tokenResponse := controller.UserService.Login(ctx, &request)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   tokenResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserControllerImpl) Get(ctx *fiber.Ctx) error {
	// Get user from context (should be set by auth middleware)
	newUser := ctx.Locals("user").(*domain.User)

	userResponse := controller.UserService.Get(ctx, *newUser)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserControllerImpl) Logout(ctx *fiber.Ctx) error {
	// Get user from context (should be set by auth middleware)
	newUser := ctx.Locals("user").(*domain.User)

	controller.UserService.Logout(ctx, *newUser)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   "Logout successful",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	// Get user from context (should be set by auth middleware)
	newUser := ctx.Locals("user").(*domain.User)

	var request user.UserUpdateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Update(ctx, *newUser, request)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
