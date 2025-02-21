package main

import (
	"context"
	"fmt"
	"github.com/dzamyatin/atomWebsite/internal/di"
	"log"
)

func main() {
	fmt.Println("Server starting...")

	ctx := context.Background()

	//server := di.InitializeGRPCServer()
	//
	//err := server.Start()

	manager := di.InitializeGRPCProcessManager()
	err := manager.Start(ctx)

	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}

	fmt.Println("Done")
}
