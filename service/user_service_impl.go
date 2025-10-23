package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
	"github.com/sorfian/go-todo-list/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (u *UserServiceImpl) Register(ctx *fiber.Ctx, request *web.UserRegisterRequest) web.TokenResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := u.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err = u.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(helper.NewNotFoundError("username already exists"))
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	token, err := helper.GenerateToken()
	helper.PanicIfError(err)

	tokenExp := helper.GetTokenExpiration(30)

	userData := domain.User{
		Username: request.Username,
		Password: hashedPassword,
		Name:     request.Name,
		Token:    token,
		TokenExp: tokenExp,
	}

	createdUser := u.UserRepository.Create(ctx, tx, userData)

	return web.TokenResponse{Token: createdUser.Token, TokenExp: createdUser.TokenExp}
}

func (u *UserServiceImpl) Login(ctx *fiber.Ctx, request *web.UserLoginRequest) web.TokenResponse {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Get(ctx *fiber.Ctx, user domain.User) web.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Logout(ctx *fiber.Ctx, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Update(ctx *fiber.Ctx, user domain.User, request web.UserUpdateRequest) web.UserResponse {
	//TODO implement me
	panic("implement me")
}
