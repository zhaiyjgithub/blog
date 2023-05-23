package resolvers

import (
	"blog/message-service/shared/helper"
	"blog/message-service/shared/model"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func SaveMessage(m model.Message) error {
	c := helper.NewDynamoDB()
	item, err := attributevalue.MarshalMap(&m)
	if err != nil {
		return err
	}
	_, err = c.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(m.TableName()),
		Item:      item,
	})
	return err
}
