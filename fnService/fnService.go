package fnService

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"os"
	"strings"
)

type Response events.APIGatewayProxyResponse

type FnRequestPayload struct {
	ResolverName string
	Body interface{}
}

type FnRequest struct {
	ServiceName string
	FunctionName string
	Payload FnRequestPayload
}

func CallFn(ctx context.Context, r FnRequest) (*lambda.InvokeOutput, error) {
	stage := os.Getenv("stage")
	var builder strings.Builder
	builder.WriteString(r.ServiceName)
	builder.WriteString("-")
	builder.WriteString(stage)
	builder.WriteString("-")
	builder.WriteString(r.FunctionName)

	fnName := builder.String()
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := lambda.NewFromConfig(cfg)

	fmt.Printf("fnName: %s\r\n", fnName)
	jb, err := json.Marshal(r.Payload)
	if err != nil {
		return nil, err
	}
	return client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String(fnName),
		Payload: jb,
	})
}
