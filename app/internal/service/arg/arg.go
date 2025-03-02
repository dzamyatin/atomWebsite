package arg

import "github.com/alexflint/go-arg"

type Arg struct {
	CommonArg
	Config string `arg:"-c,required" help:"path to config file"`
}

func MustNewArg() *Arg {
	args := &Arg{}
	err := arg.Parse(args)

	if err != nil {
		panic(err)
	}

	return args
}
