package main

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"hello! world!",
	})
}

func getFib(c *gin.Context) {
	answer := fib(400)
	c.JSON(200, gin.H{
		"result": answer,
	})
}

func fib(n int32) int32 {
	return 500
}


func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.GET("/fib", getFib)
	router.Run(":8080")
}