package email

import (
	"PluginServer/config"
	"fmt"
	"net/smtp"
)

func SendVerificationEmail(receiver string) error {
	// Sender data.
	from := config.Config.Credentials.Email

	to := []string{receiver}
	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("",
		from,
		config.Config.Credentials.EmailPassword,
		config.Config.SMTP.Host,
	)

	// Sending email.
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
	fmt.Println("Email Sent Successfully!")
	return nil
}
