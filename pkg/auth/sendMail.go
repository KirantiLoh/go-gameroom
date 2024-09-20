package auth

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(email, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("SMTP_AUTH_EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verify your Gameroom account")
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_AUTH_EMAIL"), os.Getenv("SMTP_AUTH_PASSWORD"))

	return dialer.DialAndSend(m)

}
