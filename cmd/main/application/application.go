package application

import (
	"github.com/lucaiatropulus/social/config"
	"github.com/lucaiatropulus/social/internal/auth"
	"github.com/lucaiatropulus/social/internal/mailer"
	ratelimiter "github.com/lucaiatropulus/social/internal/rate_limiter"
	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/internal/store/cache"
	"go.uber.org/zap"
)

type Application struct {
	Config        config.Config
	Store         store.Storage
	Cache         cache.Storage
	Mailer        mailer.Client
	Authenticator auth.Authenticator
	RateLimiter   ratelimiter.Limiter
}

func NewApplication(config config.Config, store store.Storage, cache cache.Storage, logger *zap.SugaredLogger, mailer mailer.Client, authenticator auth.Authenticator, rateLimiter ratelimiter.Limiter) *Application {
	return &Application{
		Config:        config,
		Store:         store,
		Cache:         cache,
		Mailer:        mailer,
		Authenticator: authenticator,
		RateLimiter:   rateLimiter,
	}
}
