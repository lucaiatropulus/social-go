package routing

import (
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/handlers"
	"github.com/lucaiatropulus/social/middleware"
)

type Routing struct {
	app        *application.Application
	handlers   *handlers.Handlers
	middleware *middleware.Middleware
}

func NewRouting(app *application.Application) *Routing {
	routingHandlers := handlers.NewHandlers(app)
	routingMiddleware := middleware.NewMiddleware(app)
	return &Routing{
		app:        app,
		handlers:   routingHandlers,
		middleware: routingMiddleware,
	}
}
