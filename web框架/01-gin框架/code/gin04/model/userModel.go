package model

type UserModel struct {
	// 二次校验邮箱的格式是否正确
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
	// 重复密码的二次校验
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}
