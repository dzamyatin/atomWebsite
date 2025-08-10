package servicemessengerdriver

type IMessengerDriver interface {
	SendMessage(message Message) error
	AskPhone(ChatLink) error
}
