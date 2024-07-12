package provider

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/daluzsi/go-message-broker/src/configuration/logger"
	"github.com/daluzsi/go-message-broker/src/configuration/properties"
)

var Config aws.Config

// InitConfig function used to initialize provider config with assume role
func InitConfig(ctx context.Context, props properties.Properties) {
	logger.Info("Initializing provider config", "InitConfig", logger.INIT)
	defer ctx.Done()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Error("Occurs an error during load aws credentials", err, "InitConfig", logger.DONE)
	}

	stsSvc := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsSvc, props.AWS.IAM.RoleArn)
	cfg.Credentials = aws.NewCredentialsCache(creds)

	Config = cfg

	logger.Debug(fmt.Sprintf("Credentials loaded: %+v", cfg), "InitConfig", logger.PROGRESS)

	logger.Info("Successfully get provider config", "InitConfig", logger.DONE)
}
