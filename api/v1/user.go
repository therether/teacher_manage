package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"teacher2/model"
	"teacher2/utils/validator"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var u model.User
	c.ShouldBind(&u)
	var msg string
	code1, msg := validator.Validate(&u)
	if code1 != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code":    code1,
			"message": msg,
		})
		return
	}

	code, text := model.CheckUser(u.Number)
	if code == 200 {
		code1, text1 := model.AddUser(&u)
		c.JSON(http.StatusOK, gin.H{
			"code": code1,
			"text": text1,
			"data": u,
		})
	}
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
			"data": u,
		})
		return
	}
}

// DeleteUser 删除用户信息
func DeleteUser(c *gin.Context) {
	var u model.User
	c.ShouldBind(&u)

	number := c.Query("number")

	code1, _ := model.CheckUser(number)
	if code1 == 200 {
		c.JSON(http.StatusOK, gin.H{
			"code1": 500,
			"text1": "用户已不存在",
		})
		return
	}

	code, text := model.DeleteUser(number)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
	})
}

// UpdateUser 教师修改自己信息
func UpdateUser(c *gin.Context) {
	var u model.User
	c.ShouldBind(&u)

	/*body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err1:", err)
		return
	}
	var u model.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println("err2:", err)
		return
	}

	fmt.Println("u:", &u)

	name := c.PostForm("name")
	fmt.Println("name:", name)
	role, _ := strconv.Atoi(c.PostForm("role"))
	fmt.Println("role:", role)
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	address := c.PostForm("address")
	employtime := c.PostForm("employtime")
	eduction := c.PostForm("eduction")
	undergraduate := c.PostForm("undergraduate")
	graduate := c.PostForm("graduate")
	doctorate := c.PostForm("doctorate")*/

	var msg string
	code1, msg := validator.Validate(&u)
	if code1 != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code":    code1,
			"message": msg,
		})
		return
	}

	//id, _ := strconv.Atoi(c.Param("id"))
	token := c.GetHeader("token")
	code, text := model.UpdateUser(token, &u)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
			"data": u,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
		"data": u,
	})
}

// Update 管理员修改用户信息
func Update(c *gin.Context) {
	var u model.User
	c.ShouldBind(&u)

	var msg string
	code1, msg := validator.Validate(&u)
	if code1 != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code":    code1,
			"message": msg,
		})
		return
	}

	//id, _ := strconv.Atoi(c.Param("id"))
	//token := c.GetHeader("token")
	code, text := model.Update(u.Number, &u)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
			"data": u,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
		"data": u,
	})
}

// SelectUserById 教师根据token查看自己信息
func SelectUserById(c *gin.Context) {
	var u model.User
	c.ShouldBind(&u)

	token := c.GetHeader("token")
	//fmt.Println("token:", token)
	//token := c.Request.Header.Get("token")
	//number := c.Param("number")

	code, text, data := model.SelectUserById(token)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
			//"data1": data1,
			//"data2": data2,
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
		//"data1": data1,
		//"data2": data2,
		"data": data,
	})
}

// SelectUserOne 管理员根据number查看用户信息
func SelectUserOne(c *gin.Context) {
	//var u model.User
	//c.ShouldBind(&u)
	//var number model.Num
	//err := c.ShouldBind(&number)
	//if err != nil {
	//	fmt.Println("err1:", err)
	//	return
	//}

	token := c.GetHeader("token")
	/*var number model.Num
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err1:", err)
		return
	}

	err = json.Unmarshal(body, &number)
	if err != nil {
		fmt.Println("err2:", err)
		return
	}*/
	number := c.Query("number")
	//fmt.Println("number:", number)

	code, text, data := model.SelectUserOne(token, number)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
		"data": data,
	})
}

// SelectAllUser 查看所有用户信息
func SelectAllUser(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}

	//data, code, text := model.SelectAllUser(pageNum, pageSize)
	code, text, data := model.SelectAllUser(pageNum, pageSize)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
			"data": data,
			/*"user":    user,
			"book":    book,
			"course":  course,
			"project": project,*/
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
		"data": data,
		/*"user":    user,
		"book":    book,
		"course":  course,
		"project": project,*/
	})
}
