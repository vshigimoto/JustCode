package repository

import (
	"booking/internal/user/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetById(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string) error
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context) error
}

type Repository struct {
	UserRepository
}

type Repo struct {
	main    sql.DB
	replica sql.DB
}

func NewRepository(main *sql.DB, replica *sql.DB) *Repo {
	return &Repo{
		main:    *main,
		replica: *replica,
	}
}
