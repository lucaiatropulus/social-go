package lifecycle

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lucaiatropulus/social/config"
	"github.com/lucaiatropulus/social/routing"
	"net/http"
	"time"
)

func Mount(routing *routing.Routing) http.Handler {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(cors.Handler(config.CorsOptions))
	r.Use(chiMiddleware.Timeout(time.Second * 60))

	routing.ConfigureRouting(r)

	return r
}
