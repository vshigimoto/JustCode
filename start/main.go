package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //
	"log"
	"start/database"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(id serial PRIMARY KEY, name text, email text)")
	if err != nil {
		return
	}
	r := gin.Default()

	if err != nil {
		return
	}

	r.GET("/user/all", database.GetUsers(db))
	r.POST("/user", database.CreateUser(db))
	r.PUT("/user/:id", database.UpdateUser(db))
	r.DELETE("/user/:id", database.DeleteUser(db))

	err = r.Run(":8080")

}
