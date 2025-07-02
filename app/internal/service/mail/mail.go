package servicemail

import "context"

type IMailerService interface {
	SendMail(ctx context.Context, to, subject, body string) error
}
