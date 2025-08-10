package servicemessengerstatemachine

type StateName string

const (
	StateInitial   StateName = "initial"
	StateWaitPhone StateName = "waitphone"
)

type IState interface {
	State() StateName

	IStateActions
}
