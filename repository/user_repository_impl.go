package repository

import (
	"context"
	"database/sql"

	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *gorm.DB, username string) (*domain.User, error) {
	user := domain.User{}
	err := tx.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (*domain.User, error) {
	SQL := `SELECT id, username, password, name, token, token_exp FROM users WHERE token = ?`
	rows, err := tx.QueryContext(ctx, SQL, token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Token, &user.TokenExp)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, sql.ErrNoRows
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *domain.User) domain.User {
	SQL := `UPDATE users SET username = ?, name = ?, token = ?, token_exp = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Name, user.Token, user.TokenExp, user.ID)
	helper.PanicIfError(err)
	return *user
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.User, error) {
	SQL := `SELECT id, username, password, name, token, token_exp FROM users WHERE id = ?`
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Token, &user.TokenExp)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, sql.ErrNoRows
}
