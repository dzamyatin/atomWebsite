package entity

import (
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
)

type Sender struct {
	PhoneNumber string                               `db:"phone_number"`
	Messenger   servicemessengerdriver.MessengerType `db:"messenger"`
	Link        servicemessengerdriver.ChatLink      `db:"link"`
}
