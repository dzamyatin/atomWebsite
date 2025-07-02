package servicemail

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MailerAtLeastOneSuccess struct {
	mailers []IMailerService
	logger  *zap.Logger
}

func NewMailerAtLeastOneSuccess(
	logger *zap.Logger,
	mailers []IMailerService,
) *MailerAtLeastOneSuccess {
	return &MailerAtLeastOneSuccess{
		logger:  logger,
		mailers: mailers,
	}
}

func (r *MailerAtLeastOneSuccess) SendMail(ctx context.Context, to, subject, body string) error {
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
