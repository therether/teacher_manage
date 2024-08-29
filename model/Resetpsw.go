package model

import (
	"fmt"
	"regexp"
	"teacher2/middleware"
)

type Resetpwd struct {
	ToEmail   string `json:"toemail"`
	InputCode string `json:"inputcode"`
	NewPsw1   string `json:"newpsw1"`
	NewPsw2   string `json:"newpsw2"`
}

func ResetPsw(token string, r *Resetpwd) (int, string) {

	if r.NewPsw1 != r.NewPsw2 {
		return 500, "两次密码不一致"
	}
	if len(r.NewPsw2) < 6 || len(r.NewPsw2) > 20 {
		return 500, "密码太长或太短"
	} else {
		lowercase := regexp.MustCompile(`[a-zA-Z]`)
		specialChar := regexp.MustCompile(`[.@$!%*#_~?&^]`)

		if !lowercase.MatchString(r.NewPsw2) {
			return 500, "密码必须包含至少一个字母"
		}

		if !specialChar.MatchString(r.NewPsw2) {
			return 500, "密码必须包含至少一个特殊字符"
		}
	}
	_, _, _, number := middleware.ParseToken(token)
	sqlStr := `update user set password=? where number=?`
	_, err := db.Exec(sqlStr, EncryptPsw(r.NewPsw2), number)

	if err != nil {
		fmt.Println("err:", err)
		return 500, "重置密码失败"
	}
	return 200, "重置密码成功"
}
