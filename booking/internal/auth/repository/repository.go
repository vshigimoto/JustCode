package repository

import (
	"database/sql"
)

type UserTokenRepository interface {
	CreateUserToken()
	UpdateUserToken()
}

type Repository interface {
	UserTokenRepository
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
