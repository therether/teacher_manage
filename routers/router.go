package routers

import (
	"github.com/gin-gonic/gin"
	v1 "teacher2/api/v1"
	"teacher2/middleware"
	"teacher2/utils"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//管理员
		auth.POST("user/add", v1.AddUser)                   //添加用户
		auth.DELETE("user/delete/:id", v1.DeleteUser)       //删除用户
		auth.PUT("teacher/update", v1.Update)               //管理员修改教师信息
		auth.GET("teacher/select/number", v1.SelectUserOne) //管理员根据token,number查看用户信息 查看单个用户信息
		auth.GET("user/select/all", v1.SelectAllUser)       //查看所有用户信息

		//普通用户
		auth.PUT("user/update", v1.UpdateUser)       //教师修改教师信息
		auth.GET("user/select", v1.SelectUserById)   //根据token查看用户信息 查看单个用户信息
		auth.POST("user/upload/:id", v1.UploadImg)   //上传头像
		auth.POST("user/sendcode", v1.SendCode)      //发送验证码
		auth.POST("user/resetpsw", v1.ResetPassword) //重置密码

		auth.PUT("user/updatepsw", v1.UpdatePsw) //修改密码
		auth.GET("user/logout", v1.Logout)       //退出登录
	}

	router := r.Group("api/v1")
	{
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
