package servicemessengerdriver

import (
	"context"
	"iter"
)

type IMessengerDriver interface {
	ReadMessages(ctx context.Context) iter.Seq2[Message, error]
	SendMessage(message Message) error
	AskPhone(ChatLink) error
	GetChatID(message Message) (string, error)
	GetUserPhone(message Message) (string, error)
	GetMessengerType() MessengerType
}
