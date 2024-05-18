package services

import (
	"fmt"
	"log"
	"net/smtp"
)

var emails emailsService
var Emails = &emails

type emailsService struct {
	config *EmailsConfig
}

type EmailsConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Devmode bool
}

func (e *emailsService) Init(config *EmailsConfig) {
	e.config = config
}

func (e *emailsService) sendEmails(recipients []string, subject, body string) {
	if e == nil || e.config == nil {
		log.Println("Emails service not initiated")
		return
	}

	addr := e.config.Host + ":" + e.config.Port

	auth := smtp.PlainAuth("", e.config.User, e.config.Pass, e.config.Host)

	msg1 := fmt.Sprintf("From: %s\r\n", e.config.User)
	msg3 := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	if e.config.Devmode {
		log.Printf("sending emails to %v\n%vTo: ...\n%v\n\n", recipients, msg1, msg3)
	}

	for _, to := range recipients {
		msg2 := fmt.Sprintf("To: %s\r\n", to)
		message := msg1 + msg2 + msg3

		err := smtp.SendMail(addr, auth, e.config.User, []string{to}, []byte(message))
		if err != nil {
			log.Println(err)
		}
	}
}

type MailingList struct {
	Subject       string
	GetRecipients func() ([]string, error)
	GetBody       func() (string, error)
}

func (e *emailsService) NewList(config *MailingList) func() {
	return func() {
		recipients, recipientsErr := config.GetRecipients()
		data, dataErr := config.GetBody()

		if recipientsErr == nil && dataErr == nil {
			e.sendEmails(recipients, config.Subject, data)
		}
	}
}
