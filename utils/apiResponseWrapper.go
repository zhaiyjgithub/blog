package utils

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type Code int

const (
	OK       Code = 200
	Error400 Code = 400
	Error500 Code = 500
)

const (
	Successful = "success"
	ParamErr   = "param error"
)

func Success(data interface{}, msg string) (events.APIGatewayProxyResponse, error) {
	body := Body{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	jbytes, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		StatusCode: int(OK),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jbytes),
	}, nil
}

func Fail(code Code, msg string, data interface{}) (events.APIGatewayProxyResponse, error) {
	body := Body{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
	jbytes, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		StatusCode: int(code),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jbytes),
	}, nil
}

type Body struct {
	Code Code        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

