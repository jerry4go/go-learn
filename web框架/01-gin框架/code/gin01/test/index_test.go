package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../initRouter"
	"github.com/stretchr/testify/assert"
)

// 单元测试  go test
func TestIndexGetRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello google", w.Body.String())
}
