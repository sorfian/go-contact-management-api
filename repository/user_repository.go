package repository

import (
	"context"
	"database/sql"

	"github.com/sorfian/go-todo-list/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	FindByUsername(ctx context.Context, tx *gorm.DB, username string) (*domain.User, error)
	FindByToken(ctx context.Context, tx *sql.Tx, token string) (*domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user *domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.User, error)
}
