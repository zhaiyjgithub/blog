package main

import (
	"blog/fnService"
	"blog/message-service/model"
	"blog/message-service/resolvers"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Route(r fnService.FnRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "",
	}
	var err error
	payload := r.Payload

	switch payload.ResolverName {
	case "saveMessage":
		m, ok := payload.Body.(model.Message)
		if !ok {
			fmt.Printf("body: %v", payload.Body)
			fmt.Println("parse body failed", err.Error())

		} else {
			_ = resolvers.SaveMessage(m)
		}
		return resp, nil
	default:
		resp.StatusCode = 400
		errText := fmt.Sprintf("%s is not found", r.ServiceName)
		err = errors.New(errText)
		return resp, err
	}
}

