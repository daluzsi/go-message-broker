package main

import (
	"context"
	"github.com/daluzsi/go-message-broker/src/adapter/input"
	"github.com/daluzsi/go-message-broker/src/configuration/logger"
	"github.com/daluzsi/go-message-broker/src/configuration/properties"
	"github.com/daluzsi/go-message-broker/src/configuration/provider"
	"os/signal"
	"syscall"
)

func main() {
	logger.Info("Initializing broker...", "main", logger.INIT)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// initialize properties
	props := properties.InitProperties()

	// initialize provider config
	provider.InitConfig(ctx, *props)

	// initialize sqs client
	sqsAdp := input.NewSQS(provider.Config, *props)

	// start polling
	sqsAdp.StartPolling(ctx)

	// wait until receive shutdown signal
	select {
	case <-ctx.Done():
		logger.Info("Shutting down broker gracefully...", "main", logger.DONE)
	}
}
