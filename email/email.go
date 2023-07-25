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

	message := []byte(fmt.Sprintf(
		`Subject: email verification

<a href="%s">
<button>verify email</button>
</a>`,
		config.Config.Server.URI()+"?sid="+sessionId.String(),
	))

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

	fmt.Println("Email Sent Successfully!")
	return nil
}
