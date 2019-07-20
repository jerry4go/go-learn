package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../initRouter"
	"github.com/stretchr/testify/assert"
)

// 单元测试  go test
// 通过 assert 进行断言，来判断返回状态码和返回值是否与代码中的值一致。

func TestIndexGetRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello google", w.Body.String())
}
