package service

import (
	"context"

	"github.com/sorfian/go-todo-list/model/domain"
	"github.com/sorfian/go-todo-list/model/web"
)

type UserService interface {
	Register(ctx context.Context, request *web.UserRegisterRequest)
	Login(ctx context.Context, request *web.UserLoginRequest) web.TokenResponse
	Get(ctx context.Context, user domain.User) web.UserResponse
	Logout(ctx context.Context, user domain.User)
	Update(ctx context.Context, user domain.User, request web.UserUpdateRequest) web.UserResponse
}
