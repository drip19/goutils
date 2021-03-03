package email

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

type EMessage struct {
	Subject string
	Body    string
	To      []string
}

func SendMail(eMessage EMessage, user, password string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": user,
		"pass": password,
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, err := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "NOTICE ME"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", eMessage.To...)                         //发送给多个用户
	m.SetHeader("Subject", eMessage.Subject)                  //设置邮件主题
	m.SetBody("text/html", eMessage.Body)                     //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err = d.DialAndSend(m)
	return err

}
