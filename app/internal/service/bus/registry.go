package bus

type HandlerUnit struct {
	Command ICommand
	Handler IHandler
	BusName BusName
}

type HandlerRegistry []HandlerUnit
