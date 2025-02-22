package cmd

import (
	"context"
	"fmt"
	"reflect"

	"github.com/alexflint/go-arg"
	"github.com/dzamyatin/atomWebsite/internal/di"
	arg2 "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"github.com/dzamyatin/atomWebsite/internal/service/config"
)

type IExecuter[ARG any] interface {
	Execute(ctx context.Context, u ARG) error
}

type Command[ARG any] struct {
	executer func(ctx context.Context, arg ARG) IExecuter[ARG]
	help     string
}

func NewCommand[ARG any](
	executer func(ctx context.Context, arg ARG) IExecuter[ARG],
	help string,
) *Command[ARG] {
	return &Command[ARG]{
		executer: executer,
		help:     help,
	}
}

type parsedConfigSetter interface {
	SetParsedConfig(cfg config.AppConfig)
}

func parseConfig(args any) {
	i := reflect.ValueOf(args).Elem().FieldByName("Arg").Interface()

	if i == nil {
		return
	}

	arg, ok := i.(arg2.Arg)

	if !ok {
		return
	}

	conf, err := di.CreateConfig(arg.Config)

	if err != nil {
		panic(fmt.Sprintf("error parsing common arg %v", err))
	}

	v, ok := args.(parsedConfigSetter)
	if !ok {
		return
	}

	v.SetParsedConfig(*conf)
}

func (r *Command[ARG]) Execute(ctx context.Context) int {
	var args ARG
	arg.MustParse(&args)
	parseConfig(&args)

	if err := r.executer(ctx, args).Execute(ctx, args); err != nil {
		fmt.Println("error:", err)

		return 1
	}

	return 0
}

func (r *Command[ARG]) Help() string {
	return r.help
}
