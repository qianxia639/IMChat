package email

import (
	"IMChat/utils"
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Email struct {
	Username string
	Password string
	Host     string
}

func (e *Email) sendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.Host + ":465",
	}

	msg := []byte("To: " + to + "\r\nFrom: " + e.Username + "\r\nSubject: " + subject + "\r\n\r\n" + body)

	conn, err := tls.Dial("tcp", e.Host+":465", tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, e.Host)
	if err != nil {
		return err
	}
	defer client.Quit()

	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(e.Username); err != nil {
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

func (e *Email) SendEmailCode(cache *redis.Client, email string) error {
	code := utils.RandomInt(100000, 999999)
	body := fmt.Sprintf("您的验证码为: %d, 有效时间为5分钟", code)
	err := e.sendMail(email, "验证信息", body)
	if err != nil {
		log.Err(err).Msg("邮件发送失败")
		return err
	}

	err = cache.Set(context.Background(), email, code, time.Minute*5).Err()
	if err != nil {
		log.Err(err).Msg("SendEmailCode error")
		return err
	}

	return nil
}

func VerifyEmailCode(cache *redis.Client, email, code string) (bool, error) {
	var emailCode string
	err := cache.Get(context.Background(), email).Scan(&emailCode)
	if err != nil {
		log.Err(err).Msg("VerifyEmailCode error")
		return false, err
	}

	return emailCode == code, nil
}
