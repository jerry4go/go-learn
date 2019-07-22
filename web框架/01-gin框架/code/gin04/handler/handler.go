package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// curl http://192.168.68.29:8080/

func Index(context *gin.Context) {

	context.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}
