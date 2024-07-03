package main

import (
	"context"
	"fmt"
	"github.com/daluzsi/go-message-broker/configuration/provider"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// initialize provider config
	provider.InitConfig(ctx)

	fmt.Println("Initializing broker...")

	// wait until receive shutdown signal
	select {
	case <-ctx.Done():
		fmt.Println("Shutting down broker gracefully...")
	}
}
