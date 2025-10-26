package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/helper"
	"github.com/sorfian/go-contact-management-api/model/domain"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/user"
	"github.com/sorfian/go-contact-management-api/repository"
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

func (service *UserServiceImpl) Register(ctx *fiber.Ctx, request *user.UserRegisterRequest) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(helper.NewResourceConflictError("username already exists"))
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

func (service *UserServiceImpl) Login(ctx *fiber.Ctx, request *user.UserLoginRequest) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newUser, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(helper.NewNotFoundError("username is incorrect"))
	}

	err = helper.VerifyPassword(newUser.Password, request.Password)
	if err != nil {
		panic(helper.NewNotFoundError("password is incorrect"))
	}

	token, err := helper.GenerateToken()
	helper.PanicIfError(err)

	tokenExp := helper.GetTokenExpiration(30)

	newUser.Token = token
	newUser.TokenExp = tokenExp

	updatedUser := service.UserRepository.Update(ctx, tx, newUser)

	return web.TokenResponse{
		Token:    updatedUser.Token,
		TokenExp: updatedUser.TokenExp,
	}
}

func (service *UserServiceImpl) Get(ctx *fiber.Ctx, newUser domain.User) user.UserResponse {
	return user.UserResponse{
		Username: newUser.Username,
		Name:     newUser.Name,
	}
}

func (service *UserServiceImpl) Logout(ctx *fiber.Ctx, user domain.User) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user.Token = ""
	user.TokenExp = 0

	service.UserRepository.Update(ctx, tx, &user)
}

func (service *UserServiceImpl) Update(ctx *fiber.Ctx, newUser domain.User, request user.UserUpdateRequest) user.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if request.Name != "" {
		newUser.Name = request.Name
	}

	if request.Password != "" {
		hashedPassword, err := helper.HashPassword(request.Password)
		helper.PanicIfError(err)
		newUser.Password = hashedPassword
	}

	updatedUser := service.UserRepository.Update(ctx, tx, &newUser)

	return user.UserResponse{
		Username: updatedUser.Username,
		Name:     updatedUser.Name,
	}
}
