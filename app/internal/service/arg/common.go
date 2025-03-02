package arg

type CommonArg struct {
	Command string `arg:"positional,required" help:"command to execute"`
}
