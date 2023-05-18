package main

import (
	"blog/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
	"os"
	"strconv"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("req body", request.Body, request.Body)
	var p Param
	if err := utils.CheckParams(request.Body, &p); err != nil {
		return utils.Fail(utils.Error400, err.Error(), nil)
	}

	qName := os.Getenv("sqs_sender_worker_queue_name")
	if len(qName) == 0 {
		log.Fatalln("sqs_sender_worker_queue_name is empty")
	}
	sqsService := NewSqsService(qName)
	skip := 0
	page := 10
	for {
		pageData := pagingData(p.Data, skip, page)
		if len(pageData) == 0 {
			break
		}
		var dataString []string
		for _, d := range pageData{
			jb, err := json.Marshal(d)
			if err != nil {
				fmt.Println("Marshal page data failed", err.Error())
			} else {
				dataString = append(dataString, string(jb))
			}
		}
		_, err := sqsService.SendBatchMessages(dataString)
		if err != nil {
			fmt.Println("Send batch messages failed", err.Error())
		}
		skip = skip + page
	}
	return utils.Success(nil, "Success")
}

func pagingData(input []Message, skip int, page int) []Message {
	if skip < len(input) {
		skip = len(input)
	}
	end := skip + page
	if end < len(input) {
		end = len(input)
	}
	return input[skip: end]
}

type SqsService struct {
	Client *sqs.Client
	QueueUrl *string
}

func NewSqsService(qName string) SqsService {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalln("new sqs client failed")
	}
	client := sqs.NewFromConfig(cfg)
	input := &sqs.GetQueueUrlInput{QueueName: aws.String(qName)}
	out, err := client.GetQueueUrl(context.TODO(), input)
	if err != nil {
		fmt.Println("Get queue url failed", err.Error())
	}
	return SqsService{
		QueueUrl: out.QueueUrl,
		Client: client,
	}
}

func (s *SqsService) SendBatchMessages(data []string) (*sqs.SendMessageBatchOutput, error) {
	if len(data) == 0 {
		return nil, errors.New("data body is empty")
	}
	if len(data) > 10 {
		return nil, errors.New("batch length maxim is 10 in a batch")
	}
	var entries []types.SendMessageBatchRequestEntry
	entries = make([]types.SendMessageBatchRequestEntry, len(data))

	for i, body := range data {
		entries[i] = types.SendMessageBatchRequestEntry{
			Id:          aws.String(strconv.Itoa(10 + i)),
			MessageBody: aws.String(body),
		}
	}
	return s.Client.SendMessageBatch(context.TODO(), &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: s.QueueUrl,
	})
}

type Param struct {
	Data []Message `json:"data" validate:"gt=0,dive"`
}

 type Message struct {
	OrganizationID string `json:"organizationID" validate:"required,gt=0"`
	CallbackURL string `json:"callbackURL"`
	CallbackURLParam map[string]string `json:"callbackURLParam"`
	ReceiverID string `json:"receiverID" validate:"required"`
	ReceiverName string `json:"receiverName" validate:"required"`
	Text string `json:"text" validate:"required,gt=0"`
}
