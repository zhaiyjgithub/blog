package main

import (
	"blog/fnService"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var r fnService.FnRequest
	if err := json.Unmarshal([]byte(request.Body), &r); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            err.Error(),
		}, err
	}
	return Route(r)
}

func main() {
	lambda.Start(Handler)
}
