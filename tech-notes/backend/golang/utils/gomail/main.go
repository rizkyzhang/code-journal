package main

import (
	"fmt"
	"gomail/dto"
	"gomail/util"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("godotenv.Load: %s", err)
		return
	}

	emailUtil := util.NewEmailUtil()

	err := emailUtil.Send(&dto.EmailUtilSendDTO{
		Recipients: []string{"test@email.com", "test2@email.com"},
		ReplyTo:    "root@email.com",
		Subject:    "Test",
		Body: `
    <h1>Test<h1>
    <br/>
    <p>This is a test</h1>
    `,
		Attachments: []string{"lock.jpeg"},
	})
	if err != nil {
		fmt.Printf("emailUtil.Send: %s", err)
		return
	}
}
