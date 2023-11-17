package validate

import (
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/errors"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

func IsUsername(username string) error {

	if len(username) > config.Config.Validation.MaxNameLen {
		return errors.ErrUsernameTooLong
	}
	if len(username) < config.Config.Validation.MinNameLen {
		return errors.ErrUsernameTooShort
	}

	p := bluemonday.StrictPolicy()
	if s := p.Sanitize(username); s != username {
		return errors.ErrInvalidUsername
	}

	return nil
}

func IsEmail(email string) error {

	if len(email) > config.Config.Validation.MaxEmailLen {
		return errors.ErrEmailTooLong
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.ErrInvalidEmail
	}

	domain := strings.Split(email, "@")[1]
	cli := http.Client{
		Timeout: 5 * time.Second,
	}

	ch := make(chan error, 1)
	defer close(ch)
	go func() {
		_, err = cli.Get("http://" + domain)
		ch <- err
	}()

	if <-ch != nil {
		return errors.ErrInvalidEmail
	}

	return nil
}

func IsPassword(password string) error {

	var conf = config.Config.Validation

	if len(password) > conf.MaxPassLen {
		return errors.ErrPassTooLong
	}

	if len(password) < conf.MinPassLen {
		return errors.ErrPassTooShort
	}

	p := bluemonday.StrictPolicy()
	if s := p.Sanitize(password); s != password {
		return errors.ErrInvalidPassword
	}

	return nil
}
