package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"../router"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	r = router.RouteCheck()
}

func TestIndexHtml(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
}
