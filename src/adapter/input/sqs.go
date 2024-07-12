package input

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/daluzsi/go-message-broker/src/configuration/logger"
	"github.com/daluzsi/go-message-broker/src/configuration/properties"
	"sync"
)

type SQS struct {
	client *sqs.Client
	props  properties.Properties
}

// NewSQS new SQS input
func NewSQS(config aws.Config, props properties.Properties) *SQS {
	//TODO refact to EndpointResolverV2
	client := sqs.NewFromConfig(config, func(o *sqs.Options) {
		o.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("x", "x", ""))
		o.EndpointResolver = sqs.EndpointResolverFromURL(props.AWS.SQS.QueueUrl)
	})

	return &SQS{
		client: client,
		props:  props,
	}
}

func (s *SQS) StartPolling(ctx context.Context) {
	logger.Info("Starting polling...", "StartPolling", logger.INIT)

	for {
		select {
		case <-ctx.Done():
			logger.Info("Stopping polling...", "StartPolling", logger.DONE)
			return
		default:
			resOtp, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
				QueueUrl: &s.props.AWS.SQS.QueueUrl,
			})

			if err != nil {
				logger.Error("Occurs an error during polling messages", err, "ListenMessages", logger.DONE)
			}

			if len(resOtp.Messages) == 0 {
				continue
			}

			s.processMessage(ctx, resOtp.Messages)
		}
	}
}

func (s *SQS) processMessage(ctx context.Context, messages []types.Message) {
	wg := sync.WaitGroup{}
	wg.Add(len(messages))

	for _, msg := range messages {
		go func(msg *types.Message) {
			if _, err := s.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				ReceiptHandle: msg.ReceiptHandle,
				QueueUrl:      &s.props.AWS.SQS.QueueUrl,
			}); err != nil {
				logger.Error("Occurs an error during message deletion", err, "ListenMessages", logger.DONE)
			} else {
				logger.Info(fmt.Sprintf("Message: %+v", *msg.Body), "ListenMessages", logger.PROGRESS)
			}
		}(&msg)
	}

	wg.Wait()
}
