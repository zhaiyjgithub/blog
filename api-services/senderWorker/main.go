/**
 * @author zhaiyuanji
 * @date 2022年06月20日 3:33 下午
 */
package main

import (
	"blog/fnService"
	"blog/message-service/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/gomail.v2"
	"os"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		fmt.Println("SQS record", record.Body)
		var m model.Message
		if err := json.Unmarshal([]byte(record.Body), &m); err != nil {
			fmt.Println("Parse record body failed", err.Error())
			return err
		}
		if err := sendEmail(m.ReceiverEmail, m.Subject, m.HtmlBody); err != nil {
			m.Status = model.Failed
		}

		// save message
		payload := fnService.FnRequestPayload{
			ResolverName: "saveMessage",
			Body: m,
		}
		out, err := fnService.CallFn(ctx, fnService.FnRequest{
			ServiceName: "ihms-message-service",
			FunctionName: "messageService",
			Payload: payload,
		})
		if out != nil {
			fmt.Printf("out: %v\r\n", out)
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		// save message, invoke message service to save message
		// send callback

	}
	return nil
}

func sendEmail(to string, subject string, htmlBody string) error {
	sender := os.Getenv("email_sender")
	password := os.Getenv("email_sender_app_password")
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, "Message Service", password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
