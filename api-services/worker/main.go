/**
 * @author zhaiyuanji
 * @date 2022年06月20日 3:33 下午
 */
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		fmt.Println("SQS record", record.Body)
	}
	return nil
}
