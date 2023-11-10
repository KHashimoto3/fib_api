package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/big"
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
func fib(n int) *big.Int {
	bigX := big.NewInt(0)
	bigY := big.NewInt(1)

	for i := 0; i < n; i++ {
		sum := big.NewInt(0)
		sum.Add(bigX, bigY)
		bigX = bigY
		bigY = sum
	}
	return bigX
}


func main() {
	router := setupRouter()
	router.Run(":8080")
}