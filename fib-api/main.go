package main

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"Hello World",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.Run(":8080")
}