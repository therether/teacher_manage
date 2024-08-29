package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"teacher2/model"
)

func Login(c *gin.Context) {
	var u model.User
	c.ShouldBind(&u)

	//number := c.PostForm("number")
	//password := c.PostForm("password")
	//var number string
	//var password string
	//c.ShouldBindJSON(&number)
	//c.ShouldBindJSON(&password)
	//
	//body, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	fmt.Println("err1:", err)
	//	return
	//}
	//u1 := make(map[string]string)
	//u1["number"] = number
	//u1["password"] = password
	//err2 := json.Unmarshal(body, &u1)
	//if err2 != nil {
	//	fmt.Println("err2:", err)
	//	return
	//}
	fmt.Println("number:", u.Number)
	fmt.Println("password:", u.Password)

	/*code1, text1 := model.CheckUser(number)
	if code1 == 1002 {
		c.JSON(http.StatusOK, gin.H{
			"code1": code1,
			"text1": text1,
		})
		return
	}*/

	code, text, data := model.CheckLogin(u.Number, u.Password)

	if text != "登录成功" || code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": text,
		})
		c.Abort()
		return
	} else {
		//token, _, _ := middleware.SetToken(number, password)
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": "登录成功",
			"data":    data,
		})
	}
}
