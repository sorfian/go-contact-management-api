package repository

import (
	"context"
	"database/sql"

	"github.com/sorfian/go-todo-list/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user *domain.User) domain.User {
	panic("implement me")
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *domain.User) domain.User {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *domain.User) {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
