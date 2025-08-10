package servicemessengerstatemachine

type StateName string

const (
	StateInitial   StateName = "initial"
	StateWaitPhone StateName = "waitphone"
)

type IState interface {
	State() StateName
	Move(machine IStateMachine, to StateName) error

	IStateActions
}

type BaseState struct{}

func (r *BaseState) Move(machine IStateMachine, to StateName) error {
	return machine.Move(to)
}
