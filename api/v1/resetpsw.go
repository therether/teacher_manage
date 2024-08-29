package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"teacher2/model"
)

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	token := c.GetHeader("token")
	/*toemail := c.PostForm("toemail")
	inputCode := c.PostForm("inputcode")
	newPsw1 := c.PostForm("newpsw1")
	newPsw2 := c.PostForm("newpsw2")*/

	var r model.Resetpwd
	c.ShouldBind(&r)

	fmt.Println("toemail:", r.ToEmail)
	fmt.Println("inputcode:", r.InputCode)
	fmt.Println("newPsw1:", r.NewPsw1)
	fmt.Println("newPsw2:", r.NewPsw2)

	storeCode := Vcode
	fmt.Println("Vcode:", Vcode)
	if r.InputCode == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "验证码输入为空",
		})
		c.Abort()
		return
	}
	if r.InputCode != storeCode {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "验证码输入错误",
		})
		c.Abort()
		return
	}
	if r.InputCode == storeCode {
		code, message := model.ResetPsw(token, &r)
		if code == 500 {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"email":    r.ToEmail,
			"code":     code,
			"message":  message,
			"password": r.NewPsw2,
		})
	}
}
