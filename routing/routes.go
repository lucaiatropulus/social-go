package routing

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (routing *Routing) ConfigureRouting(r *chi.Mux) {
	r.Use(routing.middleware.RateLimiterMiddleWare)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", routing.handlers.HealthHandler.HealthCheckHandler)

		docsURL := fmt.Sprintf("%s/doc.json", routing.app.Config.APP.Address)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		routing.configurePostsRouting(r)
		routing.configureUsersRoutes(r)
		routing.configureAuthenticationRoutes(r)
	})
}
