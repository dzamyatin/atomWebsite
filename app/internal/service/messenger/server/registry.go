package messengerserver

import (
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/pkg/errors"
)

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

func (r *MessengerServerRegistry) Get(
	t servicemessengerdriver.MessengerType,
) (servicemessengerdriver.IMessengerDriver, error) {
	if v, ok := r.items[t]; ok {
		return v, nil
	}

	variants := make([]servicemessengerdriver.MessengerType, len(r.items))
	for m := range r.items {
		variants = append(variants, m)
	}

	return nil, errors.Errorf("not found messenger \"%s\" available variants: %s", t)
}
