package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func helloGin(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method\n")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// router.GET("/", helloGin)
	// router.POST("/", helloGin)
	// router.PUT("/", helloGin)
	// router.DELETE("/", helloGin)
	// router.PATCH("/", helloGin)
	// router.HEAD("/", helloGin)
	// router.OPTIONS("/", helloGin)

	// 路由分组
	index := router.Group("/")
	{
		index.Any("", helloGin)
	}
	return router
}
