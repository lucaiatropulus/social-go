package mailer

import (
	"bytes"
	"embed"
	"html/template"
)

const (
	FromName            = "IatropuluSocial"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile string, username string, email string, data any, isSandbox bool) error
}

func parseTemplate(templateFile string, subject *bytes.Buffer, body *bytes.Buffer, data any) error {
	template, err := template.ParseFS(FS, "templates/"+templateFile)

	if err != nil {
		return err
	}

	err = template.ExecuteTemplate(subject, "subject", data)

	if err != nil {
		return err
	}

	err = template.ExecuteTemplate(body, "body", data)

	if err != nil {
		return err
	}

	return nil
}
