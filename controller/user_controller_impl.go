package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
	"github.com/sorfian/go-todo-list/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
func (controller *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	request := web.UserRegisterRequest{}
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
	request := web.UserLoginRequest{}
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
	user := ctx.Locals("user").(*domain.User)

	userResponse := controller.UserService.Get(ctx, *user)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserControllerImpl) Logout(ctx *fiber.Ctx) error {
	// Get user from context (should be set by auth middleware)
	user := ctx.Locals("user").(*domain.User)

	controller.UserService.Logout(ctx, *user)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   "Logout successful",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	// Get user from context (should be set by auth middleware)
	user := ctx.Locals("user").(*domain.User)

	var request web.UserUpdateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Update(ctx, *user, request)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
