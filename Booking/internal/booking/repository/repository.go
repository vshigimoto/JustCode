package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type BookingRepository interface {
	CreateApartment() gin.HandlerFunc
	GetApartments() gin.HandlerFunc
	UpdateApartment() gin.HandlerFunc
	DeleteApartment() gin.HandlerFunc
	GetApartmentById() gin.HandlerFunc
}

type Repository interface {
	BookingRepository
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
