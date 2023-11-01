package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //
	"log"
	"net/http"
	"start/database"
)

func tockenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const token = "tokenXcxzcasdKLDSAdxc"
		tockenHeader := ctx.GetHeader("Authorization")
		if token != tockenHeader {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

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
		v1.Use(gin.CustomRecovery(handleRecovery)) // error validation
		v1.Use(tockenAuth())
		v1.GET("/user/all", database.GetUsers(db))
		v1.POST("/user", database.CreateUser(db))
		v1.PUT("/user/:id", database.UpdateUser(db))
		v1.DELETE("/user/:id", database.DeleteUser(db))
		v1.POST("/login/:id", database.Login(db))
	}

	// Listen and serve on localhost:8080
	err = r.Run(":8080")

}
