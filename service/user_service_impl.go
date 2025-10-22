package service

import (
	"context"

	"github.com/go-playground/validator/v10"
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

func (u *UserServiceImpl) Register(ctx context.Context, request *web.UserRegisterRequest) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Login(ctx context.Context, request *web.UserLoginRequest) web.TokenResponse {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Get(ctx context.Context, user domain.User) web.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Logout(ctx context.Context, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) Update(ctx context.Context, user domain.User, request web.UserUpdateRequest) web.UserResponse {
	//TODO implement me
	panic("implement me")
}
