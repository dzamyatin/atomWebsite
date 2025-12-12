package servicemail

import "context"

type IMailer interface {
	SendMail(ctx context.Context, to, subject, body string) error
}
