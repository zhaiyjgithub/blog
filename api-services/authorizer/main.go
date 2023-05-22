/**
 * @author zhaiyuanji
 * @date 2022年06月24日 10:08 下午
 */
package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(_ context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	headers := event.Headers
	key := headers["x-api-key"]
	appName := headers["x-app-name"]
	fmt.Println("X-Api-Key", key, "X-App-Name", appName)
	ok := verifyApiKey(appName, key)
	authResponse := events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: ok}
	authResponse.Context = map[string]interface{}{
		"stringKey":  "serverless custom authorizer",
		"numberKey":  2023,
		"booleanKey": true,
	}
	return authResponse, nil
}

func verifyApiKey(appName string, apiKey string) bool {
	sign, err := calcApiKey(appName)
	if err != nil {
		return false
	}
	return sign == apiKey
}

func calcApiKey(appName string) (string, error) {
	salt := os.Getenv("salt")
	if len(salt) == 0 {
		return "", errors.New("salt is empty")
	}
	h := md5.New()
	h.Write([]byte(appName))
	h.Write([]byte(salt))
	s := h.Sum(nil)
	return hex.EncodeToString(s), nil
}
