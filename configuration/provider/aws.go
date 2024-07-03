package provider

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var Config aws.Config

const ROLENAME = "arn:aws:iam::000000000000:role/localstack-role"

// InitConfig function used to initialize provider config with assume role
func InitConfig(ctx context.Context) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	stsSvc := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsSvc, ROLENAME)

	cfg.Credentials = aws.NewCredentialsCache(creds)
	Config = cfg

	fmt.Println(fmt.Printf("%s", cfg.Credentials))
}
