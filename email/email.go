package email

import (
	"PluginServer/config"
	"fmt"
	"github.com/google/uuid"
	"net/smtp"
)

func SendVerificationEmail(receiver string, sessionId uuid.UUID) error {

	from := config.Config.Credentials.Email
	to := []string{receiver}

	subject := "Subject: Email verification for Encedeus\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`<html><body><a href="%s"><button>verify email</button></a></body></html>`, config.Config.Server.URI()+"?sid="+sessionId.String())
	message := []byte(subject + mime + body)

	// auth
	auth := smtp.PlainAuth("",
		from,
		config.Config.Credentials.EmailPassword,
		config.Config.SMTP.Host,
	)

	// sending
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Config.SMTP.Host, config.Config.SMTP.Port),
		auth,
		from,
		to,
		message,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
