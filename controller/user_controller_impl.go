package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
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
	var request web.UserRegisterRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	tokenResponse := controller.UserService.Register(ctx, &request)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   tokenResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Get(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Logout(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
