package middleware

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/lucaiatropulus/social/responses"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type userIDKey string

const userIDCtx userIDKey = "userID"

func (m *Middleware) UserIDPathParamMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := checkUserParam(w, r, m.app.Logger)

		if !ok {
			return
		}

		ctx := context.WithValue(r.Context(), userIDCtx, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromPathParameter(r *http.Request) int64 {
	user, _ := r.Context().Value(userIDCtx).(int64)
	return user
}

func checkUserParam(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) (int64, bool) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		responses.BadRequest(w, r, err, "We were unable to find the user", logger)
		return 0, false
	}

	return userID, true
}
