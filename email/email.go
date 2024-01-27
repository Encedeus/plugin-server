package email

import (
	"fmt"
	"github.com/Encedeus/pluginServer/config"
	"net/smtp"
)

func SendVerificationEmail(receiver string, sessionId string) error {

	from := config.Config.SMTP.Address
	to := []string{receiver}

	message := []byte(fmt.Sprintf(
		`Subject: email verification

<a href="%s">
<button>verify email</button>
</a>`,
		config.Config.Server.URI()+"/auth/email/verify/"+sessionId,
	))

	// auth
	auth := smtp.PlainAuth("",
		from,
		config.Config.SMTP.Password,
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
		return err
	}

	fmt.Println("Email Sent Successfully!")
	return nil
}
