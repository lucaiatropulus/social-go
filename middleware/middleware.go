package middleware

import (
	"github.com/lucaiatropulus/social/cmd/main/application"
)

type Middleware struct {
	app *application.Application
}

func NewMiddleware(app *application.Application) *Middleware {
	return &Middleware{app}
}
