package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	param int
	code int
	result string
}

func TestFibRoute(t *testing.T) {

	//テストケース追加の場合はここに追加する
	testCase := []*TestCase{
		{3, 200, "{\"result\":2}"},
		{1, 200, "{\"result\":1}"},
	}

	for _, v := range testCase {
		router := setupRouter()
		w := httptest.NewRecorder()
		request := "/fib?n=" + strconv.Itoa(v.param)
		req, _ := http.NewRequest("GET", request, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, v.code, w.Code)
		assert.Equal(t, w.Body.String(), v.result)
	}
}