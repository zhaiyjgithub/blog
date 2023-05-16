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
	fmt.Println("req body", request.Body, request.Headers)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, send message",
	}, nil
}
