package usecaseuser

import servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"

type SavePhoneInput struct {
	Source      servicemessengerdriver.MessengerType
	PhoneNumber string
	ChatLink    servicemessengerdriver.ChatLink
}

type SavePhoneUseCase struct {
}
