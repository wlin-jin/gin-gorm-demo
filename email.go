package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type MailInfo struct {
	SendUser    []string `json:"send_user"`
	CopyUser    []string `json:"copy_user"`
	Subject     string   `json:"subject"`
	HtmlContent string   `json:"html_content"`
}

func SendToMails(c *gin.Context) {
	var mail_info MailInfo
	c.BindJSON(&mail_info)

	err := SendMails(mail_info.SendUser, mail_info.CopyUser, mail_info.Subject, mail_info.HtmlContent)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": 2, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": 0, "message": "success"})
}

// 发送邮件功能待定
func SendMails(send_user []string, copy_user []string, subject string, html_content string) (err error) {
	if len(send_user) == 0 || len(copy_user) == 0 {
		return
	}
	//定义服务信息
	mailConn := map[string]string{
		"user": "",
		"pass": "",
		"host": "",
		"port": "",
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", send_user...)
	m.SetHeader("Cc", copy_user...)
	//m.SetAddressHeader("Cc", copy_user, "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html_content)

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	// Send the email to Bob, Cora and Dan.
	if err = d.DialAndSend(m); err != nil {
		panic(err)
	}
	return err
}
