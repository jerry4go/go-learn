package router

import (
	"../handler"
	"github.com/gin-gonic/gin"
)

func RouteCheck() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:name", handler.UserSave)
	router.GET("/user", handler.UserSaveByQuery)
	return router
}
