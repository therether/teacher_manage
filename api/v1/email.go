package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"teacher2/model"
	"teacher2/utils"
)

var Vcode = utils.RandomNumber(6)

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	//toemail := c.PostForm("email")

	var e model.ToEmail
	c.ShouldBind(&e)

	//fmt.Printf("fromemail:%s\n", e.FromEmail)
	fmt.Printf("toemail:%s\n", e.ToEmail)

	err := model.Email(&e, Vcode)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "发送验证码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"验证码":     Vcode,
		"code":    200,
		"message": "发送验证码成功",
	})
}

// ResetPassword 重置密码
/*func ResetPassword(c *gin.Context) {
	token := c.GetHeader("token")
	toemail := c.PostForm("toemail")
	inputCode := c.PostForm("inputcode")
	newPsw1 := c.PostForm("newpsw1")
	newPsw2 := c.PostForm("newpsw2")
	fmt.Println("toemail:", toemail)
	fmt.Println("inputcode:", inputCode)
	fmt.Println("newPsw1:", newPsw1)
	fmt.Println("newPsw2:", newPsw2)

	storeCode := Vcode
	fmt.Println("Vcode:", Vcode)
	if inputCode != storeCode {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "验证码输入错误",
		})
		c.Abort()
		return
	}
	if inputCode == storeCode {
		code, message := model.ResetPsw(token, newPsw1, newPsw2)
		if code == 500 {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":    toemail,
			"code":     code,
			"message":  message,
			"password": newPsw2,
		})
	}
}*/
