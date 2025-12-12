package servicemail

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestSmtpMailService_SendMail(t *testing.T) {
	ctx := context.Background()

	//https://mail.yandex.ru/?dpda=yes&uid=525211332#setup/client
	mailer := NewMailerSmtp(
		"smtp.yandex.com",
		465,
		"dvzamiatin@yandex.ru",
		"ugiriayxnmundnex",
		"localhost",
		zap.NewNop(),
		"dvzamiatin@yandex.ru",
		true,
	)

	err := mailer.SendMail(
		ctx,
		"zamyatin.daniil@gmail.com",
		"Test 2",
		"Hello World",
	)

	require.NoError(t, err)
}

func TestGomailSmtpMailService_SendMail(t *testing.T) {
	ctx := context.Background()

	mailer := NewMailerGomailSmtp(
		"smtp.yandex.com",
		465,
		"dvzamiatin@yandex.ru",
		"ugiriayxnmundnex",
		zap.NewNop(),
		"dvzamiatin@yandex.ru",
		true,
		1*time.Second,
	)

	err := mailer.SendMail(
		ctx,
		"zamyatin.daniil@gmail.com",
		"Test 3",
		"Hello World",
	)

	require.NoError(t, err)
}
