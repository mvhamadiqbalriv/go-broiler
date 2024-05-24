package service

import (
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"

	"gopkg.in/gomail.v2"
)

type MailServiceImpl struct {
	SenderEmail  string
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func NewMailService(senderEmail string, smtpHost string, smtpPort int, smtpUsername string, smtpPassword string) MailService {
	return &MailServiceImpl{
		SenderEmail:  senderEmail,
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
	}
}

func (m *MailServiceImpl) SendMail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", m.SenderEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.SMTPHost, m.SMTPPort, m.SMTPUsername, m.SMTPPassword)

	// Send the email
	err := dialer.DialAndSend(mailer)
	helper.PanicIfError(err)

	return nil
}