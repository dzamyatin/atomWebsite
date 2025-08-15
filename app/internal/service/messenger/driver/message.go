package servicemessengerdriver

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type MessengerType string

const (
	MessengerTypeTelegram MessengerType = "telegram"
)

type ChatLink struct {
	Telegram struct {
		ChatID int64
	}
}

func (d ChatLink) Value() (driver.Value, error) {
	res, err := json.Marshal(d)
	return res, errors.Wrap(err, "failed to marshal")
}

func (d *ChatLink) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return errors.Wrap(json.Unmarshal(v, d), "can't unmarshal JSON")
	default:
		return fmt.Errorf("cannot sql.Scan() from: %#v", v)
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
