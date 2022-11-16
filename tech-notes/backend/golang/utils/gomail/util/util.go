package util

import (
	"os"
	"strconv"

	"gomail/dto"

	"gopkg.in/gomail.v2"
)

type EmailUtil interface {
	Send(dto *dto.EmailUtilSendDTO) error
}

type baseEmailUtil struct {
}

func NewEmailUtil() EmailUtil {
	return &baseEmailUtil{}
}

func (*baseEmailUtil) Send(dto *dto.EmailUtilSendDTO) error {
	mail := gomail.NewMessage()
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	senderName := os.Getenv("SENDER_NAME")

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	if dto.ReplyTo != "" {
		mail.SetHeader("Reply-To", dto.ReplyTo)
	}

	mail.SetHeader("From", mail.FormatAddress(senderEmail, senderName))
	mail.SetHeader("To", dto.Recipients...)
	mail.SetHeader("Subject", dto.Subject)
	mail.SetBody("text/html", dto.Body)

	for _, a := range dto.Attachments {
		mail.Attach(a)
	}
	defer func() {
		for _, a := range dto.Attachments {
			os.Remove(a)
		}
	}()

	dialer := gomail.NewDialer(host, intPort, senderEmail, senderPassword)

	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
