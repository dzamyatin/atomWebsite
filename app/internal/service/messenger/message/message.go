package servicemessengermessage

type MessageMeta struct {
}

type IMessage interface {
	GetPhone() string
	GetText() string
	GetMeta() MessageMeta
}
