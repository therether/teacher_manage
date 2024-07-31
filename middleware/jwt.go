package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"teacher2/utils"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

// Blacklist 黑名单
var Blacklist map[string]bool

var setToken *jwt.Token

func init() {
	Blacklist = make(map[string]bool)
}

type MyClaims struct {
	Number   string
	Password string
	jwt.StandardClaims
}

func SetToken(number string, password string) (string, int, string) {
	expireTime := time.Now().Add(time.Hour * 24 * 7) //七天
	//expireTime := time.Now().Add(time.Second * 3)
	SetClaims := MyClaims{
		number,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "teacher2",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", 1001, "token生成失败"
	}
	return token, 200, "token生成成功"
}
func ParseToken(token string) (*MyClaims, int, string, string) {
	setToken, _ = jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		number := key.Number
		return key, 200, "token验证通过", number
	} else {
		return nil, 1002, "token无效", ""
	}
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("token")
		//fmt.Println("tokenHeader:", tokenHeader)
		if tokenHeader == "" {
			code = 1003
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "请求未携带token",
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = 1004
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "请求头中auth格式错误",
			})
			c.Abort()
			return
		}
		if len(checkToken) != 1 {
			code = 1005
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "token类型错误",
			})
			c.Abort()
			return
		}

		key, _, tCode, _ := ParseToken(checkToken[0])
		if time.Now().Unix() > key.ExpiresAt {
			code = 1006
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "token已过期",
			})
			c.Abort()
			return
		}

		if Blacklist[setToken.Raw] {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "token在黑名单中",
			})
			c.Abort()
			return
		}

		if tCode != "token验证通过" {
			code = 400
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "token生成失败",
			})
			c.Abort()
			return
		}
		c.Set("number", key.Number)
		c.Next()
	}
}
