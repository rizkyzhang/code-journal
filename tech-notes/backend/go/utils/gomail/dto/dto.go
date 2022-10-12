package dto

type EmailUtilSendDTO struct {
	Recipients  []string `json:"recipients" bson:"recipients"`
	ReplyTo     string   `json:"reply_to" bson:"reply_to"`
	Subject     string   `json:"subject" bson:"subject"`
	Body        string   `json:"body" bson:"body"`
	Attachments []string `json:"attachments"`
}
