package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher-end/database"
	"publisher-end/utils"
)

// 需要传递以下参数 username password
// 传递成功后会返回 token
func HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	match, err := database.VerifyCredential(username, password)
	if !match {
		code := http.StatusBadRequest
		message := "服务器发生错误"
		if err == nil {
			code = http.StatusNotFound
			message = "用户不存在"
		}
		c.JSON(code, gin.H{
			"message": message,
		})
		return
	}
	token, e := utils.GenerateToken(username)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token产生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}


