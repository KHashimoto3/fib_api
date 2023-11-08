package main

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type Fib struct {
	Result	int32	`json:"result"`
}

func init() {
	functions.HTTP("Fib", clucFib)
}

// フィボナッチ数を返す
func clucFib(w http.ResponseWriter, r *http.Request) {
	//仮で値を代入
	fibResult := Fib {
		result: 5000
	}
	//JSONデータの書き込み
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(fibResult); err != nil {
		log.Println(err)
	}
}