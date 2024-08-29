package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"teacher2/model"
)

// UpdatePsw 修改密码
func UpdatePsw(c *gin.Context) {
	var p model.Psw
	err := c.ShouldBind(&p)
	if err != nil {
		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request body has already been read"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	/*oldPsw := c.PostForm("oldpsw")
	newPsw1 := c.PostForm("newpsw1")
	newPsw2 := c.PostForm("newpsw2")*/

	fmt.Println("oldpsw:", p.OldPsw)
	fmt.Println("newPsw1:", p.NewPsw1)
	fmt.Println("newPsw2:", p.NewPsw2)

	token := c.GetHeader("token")

	_, code, text := model.CheckPsw(token, &p)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"data": p,
			"code": code,
			"text": text,
		})
		return
	}
	if code == 200 {
		code1, text1 := model.UpdatePsw(token, p.NewPsw2)
		c.JSON(http.StatusOK, gin.H{
			"data": p,
			"code": code1,
			"text": text1,
		})
	}

}
