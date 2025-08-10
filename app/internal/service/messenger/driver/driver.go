package driver

type Driver interface {
	SendMessage(message IMessage) error
	AskPhone() error
	GetContactFormMessage(message IMessage) (ContactFormMessage, bool, error)
}

type ContactFormMessage struct {
	Username string
	Phone    string
}
