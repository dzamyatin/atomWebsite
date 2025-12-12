package servicemail

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
	"time"
)

// https://github.com/wneessen/go-mail/wiki/Simple-Mailer-Example
type MailerGomailSmtp struct {
	logger *zap.Logger
	sender string
	client *mail.Client
}

func NewMailerGomailSmtp(
	host string,
	port uint32,
	username string,
	password string,
	logger *zap.Logger,
	sender string,
	ssl bool,
	timeout time.Duration,
) *MailerGomailSmtp {
	options := make([]mail.Option, 0)

	options = append(
		options,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithPort(int(port)),
		mail.WithTimeout(timeout),
	)

	if ssl {
		options = append(options, mail.WithSSL())
	}

	client, err := mail.NewClient(
		host,
		options...,
	)

	if err != nil {
		panic(err)
	}

	return &MailerGomailSmtp{
		logger: logger,
		sender: sender,
		client: client,
	}
}

func (r *MailerGomailSmtp) SendMail(ctx context.Context, to, subject, body string) error {
	message := mail.NewMsg()
	if err := message.From(r.sender); err != nil {
		return errors.Wrap(err, "failed to set FROM address")
	}
	if err := message.To(to); err != nil {
		return errors.Wrap(err, "failed to set TO address")
	}
	message.Subject(subject)
	message.SetBodyString(mail.TypeTextHTML, body)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := r.client.DialAndSendWithContext(ctx, message); err != nil {
		return errors.Wrap(err, "failed to deliver mail")
	}

	return nil
}
