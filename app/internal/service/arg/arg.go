package arg

import (
	"github.com/alexflint/go-arg"
	"github.com/dzamyatin/atomWebsite/internal/service/config"
)

type Arg struct {
	CommonArg
	Config       string `arg:"-c,required" help:"path to config file"`
	parsedConfig config.AppConfig
}

func (r *Arg) SetParsedConfig(cfg config.AppConfig) {
	r.parsedConfig = cfg
}

func (r Arg) GetParsedConfig() config.AppConfig {
	return r.parsedConfig
}

func MustNewArg() *Arg {
	args := &Arg{}
	err := arg.Parse(args)

	if err != nil {
		panic(err)
	}

	return args
}
