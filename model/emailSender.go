package model

import (
	"fmt"
	"net/smtp"
)

type MailSender struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (ms *MailSender) SendMail(to []string, subject string, body string) error {
	address := fmt.Sprintf("%s:%s", ms.Host, ms.Port)

	auth := smtp.PlainAuth("", ms.Username, ms.Password, ms.Host)

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	if err := smtp.SendMail(address, auth, ms.Username, to, msg); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
