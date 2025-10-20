package repository

import (
	"context"
	"database/sql"

	"github.com/sorfian/go-todo-list/helper"
	"github.com/sorfian/go-todo-list/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user *domain.User) domain.User {
	SQL := `INSERT INTO users (username, password, name, token, token_exp) VALUES (?, ?, ?, ?, ?)`

	hashedPassword, err := helper.HashPassword(user.Password)
	helper.PanicIfError(err)

	token, err := helper.GenerateToken()
	helper.PanicIfError(err)

	tokenExp := helper.GetTokenExpiration(30)

	execContext, err := tx.ExecContext(ctx, SQL, user.Username, hashedPassword, user.Name, token, tokenExp)
	helper.PanicIfError(err)

	id, err := execContext.LastInsertId()
	helper.PanicIfError(err)

	user.ID = int(id)
	user.Password = hashedPassword
	user.Token = token
	user.TokenExp = tokenExp

	return *user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*domain.User, error) {
	SQL := `SELECT id, username, password, name, token, token_exp FROM users WHERE username = ?`
	rows, err := tx.QueryContext(ctx, SQL, username)
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

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *domain.User) {
	SQL := `DELETE FROM users WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, user.ID)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*domain.User, error) {
	SQL := `SELECT id, username, password, name, token, token_exp FROM users`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Token, &user.TokenExp)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
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
