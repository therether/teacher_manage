package model

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
)

type ToEmail struct {
	//FromEmail string `json:"fromemail"`
	ToEmail string `json:"toemail"` //收件人邮箱
}

// Email 使用qq邮箱发送验证码
func Email(e *ToEmail, vcode string) error {
	mailer := gomail.NewMessage()
	//mailer.SetHeader("From", e.FromEmail)                           //发件人邮箱
	mailer.SetHeader("From", "2425626506@qq.com")                   //发件人邮箱
	mailer.SetHeader("To", e.ToEmail)                               //收件人邮箱
	mailer.SetHeader("Subject", "验证码")                              //主题：验证码
	mailer.SetBody("text/plain", fmt.Sprintf("您的验证码是:%s\n", vcode)) //内容

	dialer := gomail.NewDialer("smtp.qq.com", 465, "2425626506@qq.com", "vjgpqeemouvjdija")

	//连接SMTP服务器并发送邮件
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}

/*func ResetPsw(token, newPsw1, newPsw2 string) (int, string) {
	_, _, _, number := middleware.ParseToken(token)
	if newPsw1 != newPsw2 {
		return 500, "两次密码不一致"
	}
	sqlStr := `update user set password=? where number=?`
	_, err := db.Exec(sqlStr, EncryptPsw(newPsw2), number)

	if err != nil {
		fmt.Println("err:", err)
		return 500, "重置密码失败"
	}
	return 200, "重置密码成功"
}
*/
