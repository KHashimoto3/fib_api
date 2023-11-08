package demoapi

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("AisatuGet", aisatuGet)
}

// こんにちは！と返すだけの関数
func aisatuGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "こんにちは！")
}