package server

import (
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

func RunServer() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:8080")
}
