package repository

import (
	"booking/internal/redis/redis"
	"booking/internal/user/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Create function add new user to DB
func (u *Repo) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		err := ctx.ShouldBindJSON(&user)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		//user.Name and user.Email is arguments for prepared statement ($1 and $2)
		err = u.main.QueryRow("insert into users(name, email, password) values($1, $2, $3) returning ID", user.Name, user.Email, string(hashedPassword)).Scan(&user.Id) // $1 and $2 is prepared statement
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": user.Id})
	}
}

func (u *Repo) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := make([]entity.User, 0)

		rows, err := u.replica.Query("SELECT * from users")
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

func (u *Repo) UpdateUser() gin.HandlerFunc {
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
		_, err := u.main.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (u *Repo) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		_, err := u.main.Exec("DELETE from users WHERE id=$1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (u *Repo) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		rdb := redis.NewRedisClient("localhost:6379", "", 0) // need to read from config
		value := redis.GetValue(rdb, id)
		if value != "" {
			ctx.JSON(http.StatusOK, value)
			return
		}
		rows, err := u.replica.Query("SELECT * FROM users WHERE id=$1", id)
		var user entity.User
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
		for rows.Next() {
			if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
		}
		err = redis.SetValue(rdb, id, user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
