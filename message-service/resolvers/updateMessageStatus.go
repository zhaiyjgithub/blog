package resolvers

import (
	"blog/message-service/shared/helper"
	"blog/message-service/shared/model"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
	"time"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func UpdateMesageStatus(organizationID string, createdAt string, status model.MessageStatus) error {
	c := helper.NewDynamoDB()
	m := model.Message{}
	updateExp := expression.Set(expression.Name("Status"), expression.Value(status))
	updatedAt := time.Now().UTC().Format(time.RFC3339Nano)
	updateExp = updateExp.Set(expression.Name("UpdatedAt"), expression.Value(updatedAt))

	expr, _ := expression.NewBuilder().WithUpdate(updateExp).Build()
	_, err := c.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(m.TableName()),
		Key: map[string]types.AttributeValue{
			"OrganizationID": &types.AttributeValueMemberS{Value: organizationID},
			"CreatedAt":      &types.AttributeValueMemberS{Value: createdAt},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ConditionExpression:       aws.String("attribute_exists(OrganizationID)"),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})
	return err
}