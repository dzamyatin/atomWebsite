package servicemessengerstatemachine

type IStateRegistry interface {
	Get(state StateName) (IState, error)
}
