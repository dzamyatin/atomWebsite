package driver

type MessengerType string

const (
	MessengerTypeTelegram MessengerType = "telegram"
)

type IMessage interface {
	GetUsername() string
	GetMessengerType() MessengerType
	GetChatLink() IChatLint
	GetText() string
}

type IChatLint interface {
	GetChatLink() ChatLint
}

type ChatLint struct {
	ChatID string
}
