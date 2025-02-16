package routing

import (
	"github.com/go-chi/chi/v5"
)

func (routing *Routing) configureAuthenticationRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", routing.handlers.AuthHandler.RegisterUserHandler)
		r.Post("/login", routing.handlers.AuthHandler.LoginUserHandler)
		r.Put("/activate/{token}", routing.handlers.AuthHandler.ActivateAccountHandler)
	})
}
