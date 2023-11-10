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

	nNum, _ := strconv.Atoi(n)

	//パラメータの値が無効な場合
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

//フィボナッチ数を求める関数。計算量を抑えるために積み上げ方式を採用。
func fib(n int) int {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		sum := x + y
		x = y
		y = sum
	}
	return x
}


func main() {
	router := setupRouter()
	router.Run(":8080")
}