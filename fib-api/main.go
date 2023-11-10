package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"hello! world!",
	})
}

func getFib(c *gin.Context) {
	n := c.Query("n")

	//パラメータがない場合
	if n == "" {
		c.JSON(400, gin.H{"code": 400, "error": "Parameter 'n' is required."})
		return
	} 

	nNum, _ := strconv.ParseInt(n, 10, 64)

	if nNum < 1 {
		c.JSON(400, gin.H{"code": 400, "error": "Parameter 'n' is too small."})
		return
	}

	answer := fib(nNum)

	c.JSON(200, gin.H{
		"result": answer,
	})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", hello)
	router.GET("/fib", getFib)
	return router
}

func fib(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return fib(n - 2) + fib(n - 1)
}


func main() {
	router := setupRouter()
	router.Run(":8080")
}