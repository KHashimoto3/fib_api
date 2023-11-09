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
	answer := fib(20)
	c.JSON(200, gin.H{
		"result": answer,
	})
}

func fib(n int32) int32 {
	if n <= 2 {
		return 1
	}
	return fib(n - 2) + fib(n - 1)
}


func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.GET("/fib", getFib)
	router.Run(":8080")
}