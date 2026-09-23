package claimCfg

import (
	"errors"
	"os"
)

type Config struct {
	HTTPPort     string
	SQSQueueURL  string
	AWSEndpoint  string
	AWSSecretKey string
	AWSAccessKey string
	AWSRegion    string
}

func Load() (Config, error) {
	cfg := Config{
		HTTPPort:     getEnv("HTTP_PORT", "5000"),
		AWSRegion:    getEnv("AWS_REGION", "us-east-1"),
		AWSEndpoint:  getEnv("AWS_ENDPOINT", ""),
		AWSSecretKey: getEnv("AWS_SECRET_KEY", ""),
		AWSAccessKey: getEnv("AWS_ACCESS_KEY", ""),
		SQSQueueURL:  getEnv("SQS_QUEUE_URL", ""),
	}
	if cfg.SQSQueueURL == "" {
		return Config{}, errors.New("SQS Queue is required")
	}
	if cfg.AWSEndpoint == "" {
		return Config{}, errors.New("AWS Endpoint is required")
	}
	if cfg.AWSAccessKey == "" {
		return Config{}, errors.New("AWS Access Key is required")
	}
	if cfg.AWSSecretKey == "" {
		return Config{}, errors.New("AWS Secret Key is required")
	}

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}
	return fallback
}
