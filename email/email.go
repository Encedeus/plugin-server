package email

import (
	"fmt"
	"github.com/Encedeus/pluginServer/config"
	"net/smtp"
)

func SendVerificationEmail(receiver string, sessionId string) error {

	from := config.Config.SMTP.Address
	to := []string{receiver}

	message := fmt.Sprintf(
		`

<a href="%s">
<button>verify email</button>
</a>`,
		config.Config.Server.URI()+"/auth/email/verify/"+sessionId,
	)

	// auth
	auth := smtp.PlainAuth("",
		from,
		config.Config.SMTP.Password,
		config.Config.SMTP.Host,
	)

	// sending
	subject := "Subject: email verification\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Config.SMTP.Host, config.Config.SMTP.Port),
		auth,
		from,
		to,
		[]byte(subject+mime+message),
	)

	if err != nil {
		return err
	}

	fmt.Println("Email Sent Successfully!")
	return nil
}
