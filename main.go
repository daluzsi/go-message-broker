package main

import (
	"context"
	"fmt"
	"github.com/daluzsi/go-message-broker/src/adapter/input"
	"github.com/daluzsi/go-message-broker/src/configuration/logger"
	"github.com/daluzsi/go-message-broker/src/configuration/properties"
	"github.com/daluzsi/go-message-broker/src/configuration/provider"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Info("Initializing broker...", "main", logger.INIT)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV)

	// initialize properties
	props := properties.InitProperties()

	// initialize provider config
	provider.InitConfig(ctx, *props)

	// initialize sqs client
	sqsAdp := input.NewSQS(provider.Config, *props)

	done := make(chan bool, 1)

	for _, queue := range props.AWS.SQS.QueuesUrl {
		if !sqsAdp.QueueExists(ctx, queue) {
			logger.Error(fmt.Sprintf("Queue %s not exists", queue), nil, "main", logger.DONE)
			close(done)
		}
	}

	// start polling
	go func() {
		sqsAdp.StartPolling(ctx, done)
	}()

	// wait until receive shutdown signal or done for someone else error
	select {
	case <-stop:
		logger.Info("Shutting down broker gracefully...", "main", logger.DONE)
	case <-done:
		logger.Info("Application done", "main", logger.DONE)
	}
}
