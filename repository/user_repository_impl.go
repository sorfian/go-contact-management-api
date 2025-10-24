package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx *fiber.Ctx, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx.UserContext()).Create(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx *fiber.Ctx, tx *gorm.DB, username string) (*domain.User, error) {
	user := domain.User{}
	err := tx.WithContext(ctx.UserContext()).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepositoryImpl) FindByToken(ctx *fiber.Ctx, tx *gorm.DB, token string) (*domain.User, error) {

	user := &domain.User{}
	err := tx.WithContext(ctx.UserContext()).Where("token = ?", token).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx *fiber.Ctx, tx *gorm.DB, user *domain.User) domain.User {
	err := tx.WithContext(ctx.UserContext()).Save(user).Error
	helper.PanicIfError(err)
	return *user
}

func (repository *UserRepositoryImpl) FindById(ctx *fiber.Ctx, tx *gorm.DB, id int) (*domain.User, error) {
	user := domain.User{}
	err := tx.WithContext(ctx.UserContext()).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
