package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"start/entity"
)

// Create function add new user to DB
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

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := make([]entity.User, 0)

		rows, err := db.Query("SELECT * from users")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		defer rows.Close()
		for rows.Next() {
			var user entity.User
			if err = rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			users = append(users, user)
		}
		if err = rows.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, users)
	}

}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		var user entity.User
		// Unmarshal json to a new user
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		_, err := db.Exec("DELETE from users WHERE id=$1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
