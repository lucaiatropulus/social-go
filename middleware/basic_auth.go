package middleware

import (
	"encoding/base64"
	"github.com/lucaiatropulus/social/responses"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (m *Middleware) BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized := checkBasicAuth(w, r, m.app.Logger); !isAuthorized {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func checkBasicAuth(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) bool {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		responses.Unauthorized(w, r, logger)
		return false
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Basic" {
		responses.Unauthorized(w, r, logger)
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		responses.Unauthorized(w, r, logger)
		return false
	}

	username := "admin"
	password := "password"

	credentials := strings.SplitN(string(decoded), ":", 2)

	if len(credentials) != 2 || credentials[0] != username || credentials[1] != password {
		responses.Unauthorized(w, r, logger)
		return false
	}

	return true
}
