package messengerserver

import servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"

type MessengerServerRegistry struct {
	items map[servicemessengerdriver.MessengerType]servicemessengerdriver.IMessengerDriver
}

func NewMessengerServerRegistry(
	messengers []servicemessengerdriver.IMessengerDriver,
) *MessengerServerRegistry {
	items := make(map[servicemessengerdriver.MessengerType]servicemessengerdriver.IMessengerDriver, len(messengers))

	for _, messenger := range messengers {
		items[messenger.GetMessengerType()] = messenger
	}

	return &MessengerServerRegistry{
		items: items,
	}
}

func (r *MessengerServerRegistry) Get(t servicemessengerdriver.MessengerType) servicemessengerdriver.IMessengerDriver {
	return r.items[t]
}
