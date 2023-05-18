package model

import (
	"log"
	"os"
)

type MessageStatus string
const (
	Sending MessageStatus = "Sending"
	Sent MessageStatus = "Sent"
	Failed MessageStatus = " Failed"
)

type Message struct {
	OrganizationID string
	CreatedAt      string
	MessageID string
	CallbackURL string
	CallbackURLParam map[string]string
	ReceiverID string
	ReceiverName string
	Text string
	Status MessageStatus
}

func (m *Message) TableName() string {
	tableName := os.Getenv("message_table")
	if len(tableName) == 0 {
		log.Fatalln("message_table is empty")
	}
	return tableName
}
