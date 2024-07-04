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
	SQS SQS `yaml:"sqs"`
}

type SQS struct {
	QueueUrl string `yaml:"queue_url"`
}

func InitProperties() *Properties {
	logger.Info("Initializing properties", "InitProperties", logger.INIT)

	buf, err := os.ReadFile("application-properties.yaml")

	if err != nil {
		logger.Error("Occurs an error during read file", err, "InitProperties", logger.DONE)
		return nil
	}

	p := &Properties{}

	err = yaml.Unmarshal(buf, p)

	if err != nil {
		logger.Error("Occurs an error during yaml unmarshal", err, "InitProperties", logger.DONE)
		return nil
	}

	logger.Debug(fmt.Sprintf("Properties loaded: %+v", p), "InitProperties", logger.PROGRESS)

	logger.Info("Successfully load properties from file", "InitProperties", logger.DONE)

	return p
}
