package repository

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	claim_config "tlb-red.com/gopy-order-processing/internal/config"
	claim_models "tlb-red.com/gopy-order-processing/internal/models"
)

type ClaimQueue interface {
	claimQueue(ctx context.Context, claim SQSClaimQueue)
}

type SQSClaimQueue struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSClaimQueue(cfg claim_config.Config) (SQSClaimQueue, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AWSAccessKey, cfg.AWSSecretKey, "test")),
	)
	if err != nil {
		return SQSClaimQueue{}, err
	}

	query := SQSClaimQueue{
		client:   sqs.NewFromConfig(awsConfig),
		queueURL: cfg.SQSQueueURL,
	}

	return query, nil
}

func (q SQSClaimQueue) Publish(ctx context.Context, claim claim_models.Claim) error {
	body, err := json.Marshal(claim)

	if err != nil {
		return err
	}

	_, err = q.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &q.queueURL,
		MessageBody: aws.String(string(body)),
	})

	return err
}
