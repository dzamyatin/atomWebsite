package servicemail

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

// https://pkg.go.dev/gopkg.in/gomail.v2#example-package
type MailerSmtp struct {
	dialer *gomail.Dialer
	logger *zap.Logger
	sender string
}

func NewMailerSmtp(
	host string,
	port uint32,
	username string,
	password string,
	localName string,
	logger *zap.Logger,
	sender string,
	ssl bool,
) *MailerSmtp {
	d := gomail.NewDialer(
		host,
		int(port),
		username,
		password,
	)

	d.LocalName = localName
	d.SSL = ssl

	return &MailerSmtp{
		dialer: d,
		logger: logger,
		sender: sender,
	}
}

func (r *MailerSmtp) SendMail(_ context.Context, to, subject, body string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", r.sender)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	msg.SetHeader("Mime-Version", "1.0")

	err := r.dialer.DialAndSend(msg)
	if err != nil {
		return errors.Wrap(err, "send mail")
	}

	return nil
}
