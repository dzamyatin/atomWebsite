package servicemessengerstatemachine

type StateName string

const (
	StateInitial     StateName = "initial"
	StateWaitPhone   StateName = "waitphone"
	StatePhoneStored StateName = "phonestored"
)

type IState interface {
	State() StateName

	IStateActions
}
