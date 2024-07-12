package properties

import (
	"fmt"
	"github.com/daluzsi/go-message-broker/src/configuration/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Properties struct {
	AWS AWS `yaml:"aws"`
}

type AWS struct {
	IAM      IAM    `yaml:"iam"`
	SQS      SQS    `yaml:"sqs"`
	Region   string `yaml:"region"`
	Endpoint string `yaml:"endpoint"`
	IsLocal  bool   `yaml:"is_local"`
}
type IAM struct {
	RoleArn string `yaml:"role_arn"`
}

type SQS struct {
	QueueUrl string `yaml:"queue_url"`
}

func InitProperties() *Properties {
	logger.Info("Initializing properties", "InitProperties", logger.INIT)
	p := &Properties{}

	buf, err := os.ReadFile("application-properties.yaml")

	if err != nil {
		logger.Error("Occurs an error during read file", err, "InitProperties", logger.DONE)
		return p
	}

	err = yaml.Unmarshal(buf, p)

	if err != nil {
		logger.Error("Occurs an error during yaml unmarshal", err, "InitProperties", logger.DONE)
		return p
	}

	logger.Debug(fmt.Sprintf("Properties loaded: %+v", p), "InitProperties", logger.PROGRESS)

	logger.Info("Successfully load properties from file", "InitProperties", logger.DONE)

	return p
}
