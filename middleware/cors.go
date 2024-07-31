package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//cors.New(cors.Config{
		//AllowAllOrigins: true, //允许所有的跨域
		////AllowOrigins:  []string{"*"},
		//AllowMethods:     []string{"*"}, //请求方法
		//AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "AccessToken", "X-CSRF-Token"},
		//ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		//AllowCredentials: true,
		//AllowCredentials: true,
		/*AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},*/
		//MaxAge: 12 * time.Hour, //域请求的持续时间
		//})
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, "+
			"Accept, X-Requested-With, X-CSRF-Token,Authorization,token")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")
		c.Set("content-type", "application/json")
		//遇到所有的options请求都放行
		method := c.Request.Method
		if method == "OPTIONS" {
			c.JSON(http.StatusNoContent, "Options Request")
		}
		fmt.Println("method", method)
		c.Next()
	}
}

//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept, X-Requested-With, X-CSRF-Token")
//		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//		method := c.Request.Method
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusOK)
//		}
//		c.Next()
//	}
//}
