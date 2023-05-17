package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("req body", request.Body, request.Body)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, send message",
	}, nil
}

type Param struct {
	Data []Message
}

 type Message struct {
	OrganizationID string
	CreatedAt      string
	MessageID string
	CallbackURL string
	CallbackURLParam map[string]string
	ReceiverID string
	ReceiverName string
	Text string
}
