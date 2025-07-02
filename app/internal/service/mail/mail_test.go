package servicemail

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestSmtpMailService_SendMail(t *testing.T) {
	//https://mail.yandex.ru/?dpda=yes&uid=525211332#setup/client
	mailer := NewSmtpMailService(
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
		"zamyatin.daniil@gmail.com",
		"Test 2",
		"Hello World",
	)

	require.NoError(t, err)
}
