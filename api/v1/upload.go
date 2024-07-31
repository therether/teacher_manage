package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"teacher2/model"
)

func UploadImg(c *gin.Context) {

	/*id, _ := strconv.Atoi(c.Param("id"))
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"text": "获取文件失败",
			"err":  err.Error(),
			"file": file,
		})
		return
	}
	//保存文件到服务器
	filename := file.Filename
	_, _, dst := model.UploadImg(path.Join("./upload/", filename), id)

	imgFile, err := os.Open(dst)
	if err != nil {
		fmt.Println("打开图片文件失败:", err)
		return
	}
	defer imgFile.Close()

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"text": "保存文件到本地失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"text": "保存文件成功",
	})*/

	//id, _ := strconv.Atoi(c.Param("id"))
	//读取文件内容
	/*file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "读取文件失败",
			"error":   err,
		})
	}

	//保存图片到服务器指定目录
	dst := path.Join("./upload/" + file.Filename)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"msg": file,
	})*/
	/*code, text := model.UploadImg(dst, id)
	if code == 500 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"text": text,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"text": text,
	})*/

	/*var i model.User
	c.ShouldBind(&i)*/

	file, err := c.FormFile("imgurl")
	if err != nil {
		fmt.Println("err1:", err)
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	//保存文件到服务器指定目录
	uploadDir := "./upload/"
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		fmt.Println("err2:", err)
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	//id, _ := strconv.Atoi(c.Param("id"))

	token := c.GetHeader("token")

	code, text := model.UploadImg(token, filePath)
	if code == 500 {
		fmt.Println("err3:", err)
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
