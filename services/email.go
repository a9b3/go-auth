package services

import (
	"fmt"
	"net/smtp"
)

// Send will send email
func Send(from string, pass string, to string, subject string, body string) error {
	msg := fmt.Sprintf(`From: %s
	To: %s
	Subject: %s
	%s
	`, from, to, subject, body)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		return err
	}
	return nil
}
