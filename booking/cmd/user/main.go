package main

import (
	"booking/internal/user/database"
	"booking/internal/user/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //
	"log"
	"net/http"
)

//func tockenAuth() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		const token = "tokenXcxzcasdKLDSAdxc"
//		tockenHeader := ctx.GetHeader("Authorization")
//		if token != tockenHeader {
//			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
//			ctx.Abort()
//			return
//		}
//		ctx.Next()
//	}
//}

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(id serial PRIMARY KEY, name text, email text, password text)")
	if err != nil {
		return
	}
	r := gin.Default()

	handleRecovery := func(c *gin.Context, err any) {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Bad Request err: %v", err))
	}
	// Simple group: v1
	v1 := r.Group("/v1")
	{
		//v1.Use() it is Middleware init
		//v1.Use(tockenAuth())
		v1.Use(gin.CustomRecovery(handleRecovery)) // error validation
		v1.GET("/user/all", repository.GetUsers(db))
		v1.POST("/user", repository.CreateUser(db))
		v1.PUT("/user/:id", repository.UpdateUser(db))
		v1.DELETE("/user/:id", repository.DeleteUser(db))
		v1.POST("/login/:id", repository.Login(db))
		v1.GET("/user/:id", repository.GetByID(db))
	}

	//user := &entity.User{"1", "John", "john@gmail.com", "12345"}
	//repository.UserRepository.CreateUser(user)

	// Listen and serve on localhost:8080
	err = r.Run(":8080")

}
