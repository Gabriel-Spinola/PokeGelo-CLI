package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []user{
	{ID: "1", Name: "Gabriel"},
	{ID: "2", Name: "Admin"},
	{ID: "2", Name: "Lucao"},
}

func getUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, users)
}

func postUsers(ctx *gin.Context) {
	var newUser user

	// Bind json data into newUser
	if err := ctx.BindJSON(&newUser); err != nil {
		log.Fatal("Failed to bind json")

		return
	}

	users = append(users, newUser)
	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func RunServer() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}
