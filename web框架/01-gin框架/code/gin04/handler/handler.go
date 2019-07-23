package handler

import (
	"log"
	"net/http"
	"strings"

	"../model"
	"github.com/gin-gonic/gin"
)

// curl http://192.168.68.29:8080/

func Index(context *gin.Context) {

	context.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}

// curl -d 'email=1234@qq.com&password=123456&password-again=1234567' http://192.168.68.29:8080/user/register
func UserRegister(context *gin.Context) {

	// email := context.PostForm("email")
	// password := context.DefaultPostForm("password", "Wa123456")
	// passwordAgain := context.DefaultPostForm("password-again", "Wa123456")
	// println("email", email, "password", password, "password again", passwordAgain)

	// 数据模型绑定

	// var user model.UserModel
	// if err := context.ShouldBind(&user); err != nil {
	// 	println("err ->", err.Error())
	// 	return
	// }
	// println("email", user.Email, "password", user.Password, "password-again", user.PasswordAgain)

	// 日志输出和重定向

	var user model.UserModel

	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		context.Redirect(http.StatusMovedPermanently, "/")
	}

}
