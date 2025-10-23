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

func (service *UserServiceImpl) Register(ctx *fiber.Ctx, request *web.UserRegisterRequest) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByUsername(ctx, tx, request.Username)
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

	createdUser := service.UserRepository.Create(ctx, tx, userData)

	return web.TokenResponse{Token: createdUser.Token, TokenExp: createdUser.TokenExp}
}

func (service *UserServiceImpl) Login(ctx *fiber.Ctx, request *web.UserLoginRequest) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(helper.NewNotFoundError("username is incorrect"))
	}

	err = helper.VerifyPassword(user.Password, request.Password)
	if err != nil {
		panic(helper.NewNotFoundError("password is incorrect"))
	}

	token, err := helper.GenerateToken()
	helper.PanicIfError(err)

	tokenExp := helper.GetTokenExpiration(30)

	user.Token = token
	user.TokenExp = tokenExp

	updatedUser := service.UserRepository.Update(ctx, tx, user)

	return web.TokenResponse{
		Token:    updatedUser.Token,
		TokenExp: updatedUser.TokenExp,
	}
}

func (service *UserServiceImpl) Get(ctx *fiber.Ctx, user domain.User) web.UserResponse {
	return web.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}
}

func (service *UserServiceImpl) Logout(ctx *fiber.Ctx, user domain.User) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user.Token = ""
	user.TokenExp = 0

	service.UserRepository.Update(ctx, tx, &user)
}

func (service *UserServiceImpl) Update(ctx *fiber.Ctx, user domain.User, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		hashedPassword, err := helper.HashPassword(request.Password)
		helper.PanicIfError(err)
		user.Password = hashedPassword
	}

	updatedUser := service.UserRepository.Update(ctx, tx, &user)

	return web.UserResponse{
		Username: updatedUser.Username,
		Name:     updatedUser.Name,
	}
}
