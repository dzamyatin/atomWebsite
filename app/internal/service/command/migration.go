package command

import (
	"context"
	"github.com/alexflint/go-arg"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
)

func init() {
	GetRegistry().Register("migration", NewMigrationCommand())
}

type ArgMogration struct {
	mainarg.CommonArg
}

type MigrationCommand struct{}

func (G MigrationCommand) Execute(ctx context.Context) int {
	args := &ArgMogration{}
	arg.MustParse(args)

	return 0
}

func NewMigrationCommand() *MigrationCommand {
	return &MigrationCommand{}
}

func (G MigrationCommand) Help() string {
	return "There is no special requirements"
}
