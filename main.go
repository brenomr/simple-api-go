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

func validUser(user users.User, c *gin.Context) (bool, users.User) {
	// More info: https://gin-gonic.com/docs/examples/binding-and-validation

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return false, user
	}

	return true, user
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

	if valid_user, new_user := validUser(new_user, c); valid_user {
		users.Users_example = append(users.Users_example, new_user)
		c.JSON(http.StatusCreated, new_user)
	}
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// Useful information (not the best practice to use, but for now it's enough):
	// https://stackoverflow.com/questions/31080285/remove-element-by-value-in-go-list

	for index, user := range users.Users_example {
		if user.ID == id {
			users.Users_example = append(users.Users_example[:index], users.Users_example[index+1:]...)
		}
	}

	c.JSON(http.StatusNoContent, nil)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	var user_to_update users.User

	if valid_user, user_to_update := validUser(user_to_update, c); valid_user {
		for index, user := range users.Users_example {
			if user.ID == id {
				user_to_update.ID = id
				users.Users_example[index] = user_to_update
			}
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/", mainPage)
	r.GET("/users", getUsers)
	r.GET("/user/:id", getUser)
	r.POST("/user", newUser)
	r.DELETE("/user/:id", deleteUser)
	r.PUT("/user/:id", updateUser)
	r.Run("localhost:3000")
}
