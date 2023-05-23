package main

import (
	"blog/fnService"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(_ context.Context, payload fnService.FnRequestPayload) (fnService.FnResponse, error) {
	return Route(payload)
}

func main() {
	lambda.Start(Handler)
}
