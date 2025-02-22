package main

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/di"
	"github.com/dzamyatin/atomWebsite/internal/service/arg"
	"log"
)

func main() {
	ctx := context.Background()

	a := arg.NewArg()
	err := di.CreateConfig(a.Config)

	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	manager := di.InitializeGRPCProcessManager()
	err = manager.Start(ctx)

	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
