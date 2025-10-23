package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
)

type UserService interface {
	Register(ctx *fiber.Ctx, request *web.UserRegisterRequest) web.TokenResponse
	Login(ctx *fiber.Ctx, request *web.UserLoginRequest) web.TokenResponse
	Get(ctx *fiber.Ctx, user domain.User) web.UserResponse
	Logout(ctx *fiber.Ctx, user domain.User)
	Update(ctx *fiber.Ctx, user domain.User, request web.UserUpdateRequest) web.UserResponse
}
