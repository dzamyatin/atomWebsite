package arg

import "github.com/alexflint/go-arg"

type Arg struct {
	Config string `arg:"-c,required" help:"path to config file"`
}

func NewArg() *Arg {
	args := &Arg{}
	arg.MustParse(args)
	return args
}
