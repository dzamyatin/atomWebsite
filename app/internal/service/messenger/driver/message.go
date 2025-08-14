package servicemessengerdriver

type MessengerType string

const (
	MessengerTypeTelegram MessengerType = "telegram"
)

type ChatLink struct {
	Telegram struct {
		ChatID int64
	}
}

type Meta struct {
	MessageOwnerContact struct {
		Name        string
		PhoneNumber string
	}
}

type Message struct {
	Username      string
	MessengerType MessengerType
	ChatLink      ChatLink
	Text          string
	Meta          Meta
}

func NewAnswer(
	answer Message,
	text string,
) Message {
	return Message{
		Username:      answer.Username,
		MessengerType: answer.MessengerType,
		ChatLink:      answer.ChatLink,
		Text:          text,
	}
}

//type ContactFormMessage struct {
//	Username string
//	Phone    string
//}
