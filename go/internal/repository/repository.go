package repository

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	claimCfg "tlb-red.com/gopy-order-processing/internal/config"
)

type claimQueue interface {
	claimQueue(ctx context.Context, claim SQSClaimQueue)
}

type SQSClaimQueue struct {
	client   *sqs.Client
	queryURL string
}

func newSQSClaimQueue(cfg claimCfg.Config) (SQSClaimQueue, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AWSAccessKey, cfg.AWSSecretKey, "test")),
	)
	if err != nil {
		return SQSClaimQueue{}, err
	}

	query := SQSClaimQueue{
		client:   sqs.NewFromConfig(awsConfig),
		queryURL: cfg.SQSQueueURL,
	}

	return query, nil
}

func (q SQSClaimQueue) Publish(context.Context, claimQueue) error {
	err := errors.New("TODO")
	return err
}
