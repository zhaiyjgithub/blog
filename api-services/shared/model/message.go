package model

type MessageStatus string
const (
	Sending MessageStatus = "Sending"
	Sent    MessageStatus = "Sent"
	Failed  MessageStatus = " Failed"
)

type SqsMessage struct {
	OrganizationID string `json:"organizationID" validate:"required,gt=0"`
	CreatedAt string `json:"createdAt"`
	CallbackURL string `json:"callbackURL"`
	CallbackURLParam map[string]string `json:"callbackURLParam"`
	ReceiverID string `json:"receiverID" validate:"required"`
	ReceiverName string `json:"receiverName" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	HtmlBody string `json:"htmlBody" validate:"required"`
}
