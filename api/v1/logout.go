package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"teacher2/middleware"
)

// Logout 退出登录
func Logout(c *gin.Context) {
	tokenString := c.GetHeader("token")
	//if tokenString == "" {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 500,
	//		"text": "token在黑名单中",
	//	})
	//	return
	//}
	middleware.Blacklist[tokenString] = true
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"text": "退出登录成功",
	})
}
