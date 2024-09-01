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
		o.EndpointResolver = sqs.EndpointResolverFromURL(props.AWS.SQS.Endpoint)
	})

	return &SQS{
		client: client,
		props:  props,
	}
}

func (s *SQS) StartPolling(ctx context.Context, done chan bool) {
	defer close(done)
	var wg sync.WaitGroup
	wg.Add(len(s.props.AWS.SQS.QueuesUrl))

	logger.Info("Starting polling...", "StartPolling", logger.INIT)

	for _, queue := range s.props.AWS.SQS.QueuesUrl {
		go func(wg *sync.WaitGroup, queueUrl string) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					logger.Info("Stopping polling...", "StartPolling", logger.DONE)
					return
				default:
					func() {
						defer func() {
							if r := recover(); r != nil {
								logger.Warn("Recovering application after a panic", nil, "StartPolling", logger.PROGRESS)
								return
							}
						}()

						resOtp, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
							QueueUrl: &queueUrl,
						})

						if err != nil {
							logger.Error(fmt.Sprintf("Occurs an error during polling messages from queue: %s", queueUrl), err, "StartPolling", logger.DONE)
							panic(err)
						}

						if len(resOtp.Messages) == 0 {
							return
						}

						go s.processMessage(ctx, resOtp.Messages, queueUrl)
					}()
				}
			}
		}(&wg, queue)
	}

	wg.Wait()
}

func (s *SQS) processMessage(ctx context.Context, messages []types.Message, queue string) {
	wg := sync.WaitGroup{}
	wg.Add(len(messages))

	for _, msg := range messages {
		go func(msg *types.Message) {
			if _, err := s.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				ReceiptHandle: msg.ReceiptHandle,
				QueueUrl:      &queue,
			}); err != nil {
				logger.Error(fmt.Sprintf("Occurs an error during message deletion from queue: %s", queue), err, "processMessage", logger.DONE)
			} else {
				logger.Info(fmt.Sprintf("Message: %+v from queue %s", *msg.Body, queue), "processMessage", logger.PROGRESS)
			}
		}(&msg)
	}

	wg.Wait()

	logger.Debug("No more messages...", "processMessage", logger.DONE)
}

func (s *SQS) QueueExists(ctx context.Context, queue string) bool {
	_, err := s.client.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(queue),
	})

	return err == nil
}
