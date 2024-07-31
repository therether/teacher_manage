package model

import (
	"fmt"
	"teacher2/middleware"
)

/*func UploadImg(dst string, id int) (int, string) {
	sqlStr := `insert into teacher2 set imgurl=? where id=?`
	_, err := db.Exec(sqlStr, dst, id)
	if err != nil {
		return 500, "保存图片路径到数据库失败"
	}
	return 200, "保存图片路径到数据库成功"
}*/

func UploadImg(token string, url string) (int, string) {
	_, _, _, number := middleware.ParseToken(token)

	sqlStr := `update user set imgurl=? where number=?`
	_, err := db.Exec(sqlStr, url, number)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "添加到数据库失败"
	}
	return 200, "添加到数据库成功"
}
