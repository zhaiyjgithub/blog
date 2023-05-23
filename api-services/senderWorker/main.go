/**
 * @author zhaiyuanji
 * @date 2022年06月20日 3:33 下午
 */
package main

import (
	"blog/api-services/shared/model"
	"blog/fnService"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/gomail.v2"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		var sm model.SqsMessage
		if err := json.Unmarshal([]byte(record.Body), &sm); err != nil {
			fmt.Println("Unmarshal record body failed", err.Error())
			return err
		}
		if err := saveMesage(ctx, record.Body); err != nil {
			fmt.Println("Save message failed", err.Error(), record.Body)
			continue
		}

		err := sendEmail(sm.ReceiverID, sm.Subject, sm.HtmlBody)
		status := model.Sending
		if err != nil {
			status = model.Failed
		} else {
			status = model.Sent
		}
		// send email and updte message status by pk-sk
		err = updateMessageStatus(ctx, sm.OrganizationID, sm.CreatedAt, status)
		if err != nil {
			fmt.Printf("Update message status failed: %v\r\n", err.Error())
		}
		// send callback if callbackURL is NOT empty
	}
	return nil
}

func saveMesage(ctx context.Context, body string) error {
	payload := fnService.FnRequestPayload{
		ResolverName: "saveMessage",
		Body:         body,
	}
	_, err := fnService.CallFn(ctx, fnService.FnRequest{
		ServiceName:  "ihms-message-service",
		FunctionName: "messageService",
		Payload:      payload,
	})
	return err
}

func updateMessageStatus(ctxt context.Context, organzationID string, createdAt string, status model.MessageStatus) error {
	body := make(map[string]string)
	body["OrganizationID"] = organzationID
	body["CreatedAt"] = createdAt
	body["status"] = string(status)
	jb, _ := json.Marshal(body)

	payload := fnService.FnRequestPayload{
		ResolverName: "updateMessageStatus",
		Body:         string(jb),
	}
	_, err := fnService.CallFn(ctxt, fnService.FnRequest{
		ServiceName:  "ihms-message-service",
		FunctionName: "messageService",
		Payload:      payload,
	})
	return err
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
