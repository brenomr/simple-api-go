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

func newUser(c *gin.Context) {
	var new_user users.User

	// Useful information:
	// https://gin-gonic.com/docs/examples/binding-and-validation

	if err := c.BindJSON(&new_user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	users.Users_example = append(users.Users_example, new_user)

	c.JSON(http.StatusCreated, new_user)
}

func main() {
	r := gin.Default()

	r.GET("/", mainPage)
	r.GET("/users", getUsers)
	r.POST("/user", newUser)
	r.Run("localhost:3000")
}
