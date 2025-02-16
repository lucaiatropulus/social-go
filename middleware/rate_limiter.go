package middleware

import (
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
)

func (m *Middleware) RateLimiterMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limitExceeded := checkRateLimiter(w, r, m.app); limitExceeded {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func checkRateLimiter(w http.ResponseWriter, r *http.Request, app *application.Application) bool {
	if app.Config.RateLimiter.Enabled {
		if allow, retryAfter := app.RateLimiter.Allow(r.RemoteAddr); !allow {
			responses.RateLimitExceededResponse(w, r, retryAfter.String(), app.Logger)
			return true
		}
	}

	return false
}
