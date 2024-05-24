package service

type MailService interface {
	SendMail(to string, subject string, body string) error
}