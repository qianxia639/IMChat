package utils

import (
	"crypto/tls"
	"net/smtp"
)

var (
	username = "2274000859@qq.com"
	password = "sersteijzokudjfj"
	host     = "smtp.qq.com"
)

func SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", username, password, host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host + ":465",
	}

	msg := []byte("To: " + to + "\r\nFrom: " + username + "\r\nSubject: " + subject + "\r\n\r\n" + body)

	conn, err := tls.Dial("tcp", host+":465", tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer client.Quit()

	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(username); err != nil {
		return err
	}

	if err := client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
