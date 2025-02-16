package mailer

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}

func NewSendgridMailer(fromEmail string, apiKey string) *SendGridMailer {
	client := sendgrid.NewSendClient(apiKey)

	return &SendGridMailer{
		fromEmail: fromEmail,
		apiKey:    apiKey,
		client:    client,
	}
}

func (m *SendGridMailer) Send(templateFile string, username string, email string, data any, isSandbox bool) error {
	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	template, err := template.ParseFS(FS, "templates/"+templateFile)

	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)

	err = template.ExecuteTemplate(subject, "subject", data)

	if err != nil {
		return err
	}

	body := new(bytes.Buffer)

	err = template.ExecuteTemplate(body, "body", data)

	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())

	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSandbox,
		},
	})

	for i := 0; i < maxRetries; i++ {
		response, err := m.client.Send(message)
		if err != nil {
			log.Printf("Failed to send email to %v, attemt %d of %d", email, i+1, maxRetries)
			log.Printf("Error: %v", err.Error())

			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		log.Printf("Email has been sent with status code %v", response.StatusCode)
		return nil
	}

	return fmt.Errorf("failed to send email after %d attempts", maxRetries)
}
