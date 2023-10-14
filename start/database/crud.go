package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"start/entity"
)

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			return
		}
		//user.Name and user.Email is arguments for prepared statement ($1 and $2)
		err = db.QueryRow("insert into users(name, email) values($1, $2) returning ID", user.Name, user.Email).Scan(&user.Id) // $1 and $2 is prepared statement
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": user.Id})
	}
}
