package resolvers

import (
	"blog/message-service/shared/model"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	os.Setenv("message_table", "ihms-message-service-dev-message")
	m := model.Message {
		OrganizationID: "ac",
		MessageID: "12346",
		CreatedAt: time.Now().UTC().Format(time.RFC3339Nano),
		ReceiverID: "yuanji.zhai@outlook.com",
		ReceiverName: "yuanji",
		Subject: "Subject",
		HtmlBody: "Body",
		UpdatedAt: time.Now().UTC().Format(time.RFC3339Nano),
	}
	if err := SaveMessage(m); err != nil {
		t.Error(err.Error())
	}else {
		fmt.Println("Save message success")
	}
}