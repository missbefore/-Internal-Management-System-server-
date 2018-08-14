package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显示绑定声明绑定 multipart 表单：
		// c.ShouldBindWith(&form, binding.Form
		// 或者你可以使用 ShouldBind 方法去简单的使用自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择适合的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")
}