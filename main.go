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

func getUser(c *gin.Context) {
	id := c.Param("id")

	for index, user := range users.Users_example {
		if user.ID == id {
			c.JSON(http.StatusOK, &users.Users_example[index])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "The user with id " + id + " was not found.",
	})
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
	r.GET("/user/:id", getUser)
	r.POST("/user", newUser)
	r.Run("localhost:3000")
}
