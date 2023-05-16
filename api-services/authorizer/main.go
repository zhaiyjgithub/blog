/**
 * @author zhaiyuanji
 * @date 2022年06月24日 10:08 下午
 */
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(_ context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {

	authResponse := events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: true}
	authResponse.Context = map[string]interface{}{
		"stringKey":  "serverless authorizer",
		"numberKey":  2023,
		"booleanKey": true,
	}
	return authResponse, nil
}
