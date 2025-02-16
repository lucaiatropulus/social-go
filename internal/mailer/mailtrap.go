package mailer

import (
	"bytes"
	"errors"

	"gopkg.in/gomail.v2"
)

type MailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey string, fromEmail string) (MailtrapClient, error) {
	if apiKey == "" {
		return MailtrapClient{}, errors.New("main key is required")
	}

	return MailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m MailtrapClient) Send(templateFile string, uername string, email string, data any, isSandbox bool) error {
	subject := new(bytes.Buffer)
	body := new(bytes.Buffer)

	if err := parseTemplate(templateFile, subject, body, data); err != nil {
		return err
	}

	message := m.setUpMessage(email, subject, body)

	if err := m.sendMail(message); err != nil {
		return err
	}

	return nil
}

func (m MailtrapClient) setUpMessage(email string, subject *bytes.Buffer, body *bytes.Buffer) *gomail.Message {
	message := gomail.NewMessage()

	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())

	message.AddAlternative("text/html", body.String())

	return message
}

func (m MailtrapClient) sendMail(message *gomail.Message) error {
	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "main", m.apiKey)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
