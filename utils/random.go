package utils

import (
	"math/rand"
	"time"
)

// RandomNumber 生成一个六位的随机数
func RandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")
	b := make([]rune, length) //length为长度
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
