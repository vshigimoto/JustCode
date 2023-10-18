package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"start/entity"
)

// Create function add new user to DB
func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		err := ctx.ShouldBindJSON(&user)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		//user.Name and user.Email is arguments for prepared statement ($1 and $2)
		err = db.QueryRow("insert into users(name, email, password) values($1, $2, $3) returning ID", user.Name, user.Email, string(hashedPassword)).Scan(&user.Id) // $1 and $2 is prepared statement
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
			if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
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

func Login(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login entity.Login
		var user entity.User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			return
		}
		rows, err := db.Query("SELECT * FROM users WHERE id=$1", login.Id)
		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not found user"})
			return
		}
		defer rows.Close()
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "not correct password"})
			return
		}
	}
}
