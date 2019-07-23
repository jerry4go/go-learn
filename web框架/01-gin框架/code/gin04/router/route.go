package router

import (
	"../handler"
	"github.com/gin-gonic/gin"
)

func RouteCheck() *gin.Engine {
	router := gin.Default()

	//router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/index.gohtml")
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	router.Static("/statics", "./statics")
	router.StaticFile("/favicon.ico", "./statics/favicon.ico")

	router.GET("/", handler.Index)
	router.POST("/user/register", handler.UserRegister)
	return router
}
