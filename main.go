package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the main page.",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", mainPage)
	r.Run("localhost:3000")
}
