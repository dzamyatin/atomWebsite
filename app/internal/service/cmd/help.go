package cmd

import (
	"context"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/davecgh/go-spew/spew"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"sort"
	"strings"
)

func init() {
	GetRegistry().Register("help", NewHelpCommand())
	GetRegistry().SetDefault(NewHelpCommand())
}

type ArgHelp struct {
	mainarg.CommonArg
	CommandHelp string `arg:"positional" help:"command to help"`
}

type HelpCommand struct{}

func (G HelpCommand) Execute(ctx context.Context) int {
	args := &ArgHelp{}
	err := arg.Parse(args)

	if err == nil {
		if args.CommandHelp != "" {
			c, ok := GetRegistry().commands[args.CommandHelp]
			if !ok {
				fmt.Printf("command not found %s", args.CommandHelp)

				return 0
			}

			fmt.Println(c.Help())

			return 0
		}
	}

	commandNames := make([]string, 0, len(GetRegistry().commands))
	for c := range GetRegistry().commands {
		commandNames = append(commandNames, c)
	}

	sort.Strings(commandNames)

	fmt.Printf(spew.Sprintf(`
Usage:
	./app  grpc --config config-local.yaml
There is another commands: 
%s
To get help about a command run: ./app help <command>
`, strings.Join(commandNames, "\n")))

	return 0
}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (G HelpCommand) Help() string {
	return "There is no special requirements"
}
