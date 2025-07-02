package servicemail

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MailerAtLeastOneSuccess struct {
	mailers []IMailer
	logger  *zap.Logger
}

func NewMailerAtLeastOneSuccess(
	logger *zap.Logger,
	mailers []IMailer,
) *MailerAtLeastOneSuccess {
	return &MailerAtLeastOneSuccess{
		logger:  logger,
		mailers: mailers,
	}
}

func (r *MailerAtLeastOneSuccess) SendMail(ctx context.Context, to, subject, body string) error {
	if len(r.mailers) == 0 {
		return errors.New("mailers is empty")
	}

	var err error
	for _, mailer := range r.mailers {
		err = mailer.SendMail(ctx, to, subject, body)
		if err == nil {
			return nil
		}
	}

	if err != nil {
		return errors.Wrap(err, "send mail at least one success")
	}

	return nil
}
