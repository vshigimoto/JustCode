package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //
	"log"
	"start/database"
	"start/logger"
)

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

	// Simple group: v1
	v1 := r.Group("/v1")
	{
		//v1.Use() it is Middleware init
		v1.Use(logger.Logger()) // write logger in future
		v1.GET("/user/all", database.GetUsers(db))
		v1.POST("/user", database.CreateUser(db))
		v1.PUT("/user/:id", database.UpdateUser(db))
		v1.DELETE("/user/:id", database.DeleteUser(db))
	}

	// Listen and serve on localhost:8080
	err = r.Run(":8080")

}
