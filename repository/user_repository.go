package repository

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-todo-list/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx *fiber.Ctx, tx *gorm.DB, user domain.User) domain.User
	FindByUsername(ctx *fiber.Ctx, tx *gorm.DB, username string) (*domain.User, error)
	FindByToken(ctx *fiber.Ctx, tx *sql.Tx, token string) (*domain.User, error)
	Update(ctx *fiber.Ctx, tx *sql.Tx, user *domain.User) domain.User
	FindById(ctx *fiber.Ctx, tx *sql.Tx, id int) (*domain.User, error)
}
