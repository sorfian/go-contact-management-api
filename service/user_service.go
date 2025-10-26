package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/user"
)

type UserService interface {
	Register(ctx *fiber.Ctx, request *user.UserRegisterRequest) web.TokenResponse
	Login(ctx *fiber.Ctx, request *user.UserLoginRequest) web.TokenResponse
	Get(ctx *fiber.Ctx, user domain.User) user.UserResponse
	Logout(ctx *fiber.Ctx, user domain.User)
	Update(ctx *fiber.Ctx, user domain.User, request user.UserUpdateRequest) user.UserResponse
}
