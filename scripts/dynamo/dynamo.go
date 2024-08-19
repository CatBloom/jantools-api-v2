package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDB interface {
	GetClient() *dynamodb.Client
}

type dynamoDB struct {
	dynamo *dynamodb.Client
}

func NewDynamoDB() DynamoDB {
	return &dynamoDB{connectDynamoDB()}
}

func (d *dynamoDB) GetClient() *dynamodb.Client {
	return d.dynamo
}

func connectDynamoDB() *dynamodb.Client {

	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	// if err != nil {
	// 	log.Fatalf("unable to load SDK config, %v", err)
	// }

	// Dynamo localのエンドポイント
	endpoint := "http://dynamodb-local:8000"

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("dummy", "dummy", "dummy")),
		config.WithRegion("ap-northeast-1"),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	return svc
}
