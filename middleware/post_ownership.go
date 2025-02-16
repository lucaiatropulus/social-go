package middleware

import (
	"github.com/lucaiatropulus/social/responses"
	"net/http"
)

func (m *Middleware) CheckPostOwnershipMiddleware(requiredRoleName string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetAuthenticatedUserFromCtx(r)
		post := GetPostFromPathParam(r)

		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		allowed, err := checkRolePrecedence(r.Context(), user, requiredRoleName, m.app)

		if err != nil {
			responses.InternalServerError(w, r, err, m.app.Logger)
			return
		}

		if !allowed {
			responses.Forbidden(w, r, m.app.Logger)
			return
		}

		next.ServeHTTP(w, r)
	})
}
