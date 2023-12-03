package repository

import (
	"booking/internal/booking/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Repo) CreateApartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var apartment entity.Apartment
		err := ctx.ShouldBindJSON(&apartment)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		err = r.main.QueryRow("insert into apartment(phone, address,category, rating) values($1, $2, $3, $4) returning ID", apartment.Phone, apartment.Address, apartment.Category, apartment.Rating).Scan(&apartment.Id) // $1 and $2 is prepared statement
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": apartment.Id})
	}
}

func (r *Repo) GetApartments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apartments := make([]entity.Apartment, 0)

		rows, err := r.replica.Query("SELECT * from apartment")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		for rows.Next() {
			var apartment entity.Apartment
			if err = rows.Scan(&apartment.Id, &apartment.Phone, &apartment.Address, &apartment.Category, &apartment.Rating); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			apartments = append(apartments, apartment)
		}
		if err = rows.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, apartments)
	}
}

func (r *Repo) UpdateApartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		var apartment entity.Apartment
		// Unmarshal json to a new user
		if err := ctx.ShouldBindJSON(&apartment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		_, err := r.main.Exec("UPDATE apartment SET phone=$1, address=$2, category=$3, rating=$4 WHERE id=$5", apartment.Phone, apartment.Address, apartment.Category, apartment.Rating, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (r *Repo) DeleteApartment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		_, err := r.main.Exec("DELETE from apartment WHERE id=$1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (r *Repo) GetApartmentById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("id")
		if login == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		rows, err := r.replica.Query("SELECT * FROM apartment WHERE id=$1", login)
		var apartment entity.Apartment
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		rows.Next()
		if err = rows.Scan(&apartment.Id, &apartment.Phone, &apartment.Address, &apartment.Category, &apartment.Rating); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, apartment)
	}
}
