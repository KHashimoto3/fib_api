package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	param	 int	//パラメータn
	code 	 int	//返されるステータスコード
	response string	//返されるレスポンス
}

func TestFibRoute(t *testing.T) {

	//テストケース追加の場合はここに追加する
	testCase := []*TestCase{
		{3, 200, "{\"result\":2}"},
		{1, 200, "{\"result\":1}"},
		{-1, 400, "{\"code\":400,\"error\":\"Parameter 'n' is too small.\"}"},
		{50, 200, "{\"result\":12586269025}"},
		{99, 200, "{\"result\":218922995834555169026}"},
	}

	//nパラメータを渡さない場合のテスト
	router := setupRouter()
	w := httptest.NewRecorder()
	request := "/fib"
	req, _ := http.NewRequest("GET", request, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, w.Body.String(), "{\"code\":400,\"error\":\"Parameter 'n' is required.\"}")

	//値が一致するかのテスト
	for _, v := range testCase {
		router := setupRouter()
		w := httptest.NewRecorder()
		request := "/fib?n=" + strconv.Itoa(v.param)
		req, _ := http.NewRequest("GET", request, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, v.code, w.Code)
		assert.Equal(t, w.Body.String(), v.response)
	}
}