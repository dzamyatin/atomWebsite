package main

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/service/command"
	"os"
)

func main() {
	ctx := context.Background()

	com := ""
	if len(os.Args) > 1 {
		com = os.Args[1]
	}

	command.GetRegistry().MustExecuteCommand(ctx, com)
}
