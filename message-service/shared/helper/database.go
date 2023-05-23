package helper

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"sync"
)

var (
	dynamoDbOnce   sync.Once
	dynamoDbClient *dynamodb.Client
)

func NewDynamoDB() *dynamodb.Client {
	dynamoDbOnce.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic(err)
		}
		dynamoDbClient = dynamodb.NewFromConfig(cfg)
	})
	return dynamoDbClient
}
