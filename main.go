package main

import (
	"net/http"
	"simple-api-go/users"

	"github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the main page.",
	})
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users.Users_example)
}

func main() {
	r := gin.Default()

	r.GET("/", mainPage)
	r.GET("/users", getUsers)
	r.Run("localhost:3000")
}
