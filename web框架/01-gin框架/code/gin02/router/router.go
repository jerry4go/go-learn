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
	
        // 第一种写法，单独写每一个方法和对应的handler
	
	// router.GET("/", helloGin)
	// router.POST("/", helloGin)
	// router.PUT("/", helloGin)
	// router.DELETE("/", helloGin)
	// router.PATCH("/", helloGin)
	// router.HEAD("/", helloGin)
	// router.OPTIONS("/", helloGin)

	// 第二种方法 路由分组
	
	index := router.Group("/")
	{
		index.Any("", helloGin)
	}
	return router
}
