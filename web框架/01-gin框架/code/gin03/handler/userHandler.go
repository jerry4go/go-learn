package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl http://192.168.68.29:8080/user/xiaoming
func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户 "+username+" 已经保存\n")
}

// curl "http://192.168.68.29:8080/user?age=18&name=Lucy"
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	//age := context.Query("age")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户:"+username+" 年龄:"+age+" 已经保存\n")
}
