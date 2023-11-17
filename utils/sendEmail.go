package utils

import (
	"net/smtp"
	"team-work-be/config"
)

func SendEmail(emailTo string, content string) error {
	from := config.GetEmailSend()
	password := config.GetPasswordEmailSend()

	// Receiver email address.
	to := []string{emailTo}

	// smtp server configuration.
	smtpHost := config.GetSmtpHost()
	smtpPort := config.GetSmtpPort()

	// Message.
	message := []byte(content)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	return err
}
