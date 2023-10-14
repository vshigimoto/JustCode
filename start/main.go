package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //
	"log"
	"net/http"
	"start/database"
)

func home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	if err != nil {
		return
	}

	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})
	r.POST("/create-user", database.CreateUser(db))

	err = r.Run(":8080")
}
