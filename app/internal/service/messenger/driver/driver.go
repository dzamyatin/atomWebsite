package servicemessengerdriver

type IMessengerDriver interface {
	SendMessage(message Message) error
	AskPhone(ChatLink) error
	GetChatID(message Message) (string, error)
}
