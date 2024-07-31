package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"teacher2/middleware"
)

type Psw struct {
	OldPsw  string `json:"oldpsw"`
	NewPsw1 string `json:"newpsw1"`
	NewPsw2 string `json:"newpsw2"`
}

// CheckPsw 检查修改密码
func CheckPsw(token string, p *Psw) (string, int, string) {

	_, _, _, number := middleware.ParseToken(token)
	//fmt.Println("token:", token)

	//fmt.Println("p:", p)
	var password string
	sqlStr := `select password from user where number=?`
	err := db.QueryRow(sqlStr, number).Scan(&password)
	if err != nil {
		fmt.Println("err:", err)
		return "", 500, "查找密码失败"
	}
	/*if oldPsw == "" || newPsw1 == "" || newPsw2 == "" {
		return "", 500, "密码为空"
	}
	if password != oldPsw && password != EncryptPsw(oldPsw) {
		fmt.Println("password:", password)
		fmt.Println("oldPsw:", oldPsw)
		return "", 500, "旧密码输入错误"
	}
	if newPsw1 != newPsw2 {
		return "", 500, "两次密码输入不同"
	}
	if len(newPsw2) < 6 || len(newPsw2) > 20 {
		return "", 500, "密码太长或太短"
	} else {
		lowercase := regexp.MustCompile(`[a-zA-Z]`)
		specialChar := regexp.MustCompile(`[.@$!%*#_~?&^]`)

		if !lowercase.MatchString(newPsw2) {
			return "", 500, "密码必须包含至少一个字母"
		}

		if !specialChar.MatchString(newPsw2) {
			return "", 500, "密码必须包含至少一个特殊字符"
		}*/

	if p.OldPsw == "" || p.NewPsw1 == "" || p.NewPsw2 == "" {
		return "", 500, "密码为空"
	}
	if password != p.OldPsw && password != EncryptPsw(p.OldPsw) {
		fmt.Println("password:", password)
		fmt.Println("oldPsw:", p.OldPsw)
		return "", 500, "旧密码输入错误"
	}
	if p.NewPsw1 != p.NewPsw2 {
		return "", 500, "两次密码输入不同"
	}
	if len(p.NewPsw2) < 6 || len(p.NewPsw2) > 20 {
		return "", 500, "密码太长或太短"
	} else {
		lowercase := regexp.MustCompile(`[a-zA-Z]`)
		specialChar := regexp.MustCompile(`[.@$!%*#_~?&^]`)

		if !lowercase.MatchString(p.NewPsw2) {
			return "", 500, "密码必须包含至少一个字母"
		}

		if !specialChar.MatchString(p.NewPsw2) {
			return "", 500, "密码必须包含至少一个特殊字符"
		}
	}
	return password, 200, "查找密码成功"
}

// UpdatePsw 修改密码
func UpdatePsw(token string, newPsw2 string) (int, string) {

	_, _, _, number := middleware.ParseToken(token)

	sqlStr1 := `update user set password=? where number=?`
	_, err = db.Exec(sqlStr1, EncryptPsw(newPsw2), number)
	//sqlStr2 := `update admin set password=? where number=?`
	//_, err = db.Exec(sqlStr2, EncryptPsw(newPsw2), number)
	if err != nil {
		fmt.Println("err:", err)
		return 500, "修改密码失败"
	}

	return 200, "修改密码成功"
}

// EncryptPsw 密码加密
func EncryptPsw(password string) string {
	newPsw := password + "teacher2"
	hash := md5.New()
	hash.Write([]byte(newPsw))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
