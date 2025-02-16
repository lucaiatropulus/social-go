package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
	"strconv"
	"strings"
)

type authUserKey string

const authUserCtx authUserKey = "authUserKey"

func (m *Middleware) JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := checkJWTAuth(w, r, m.app)

		if user == nil {
			return
		}

		ctx := context.WithValue(r.Context(), authUserCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAuthenticatedUserFromCtx(r *http.Request) *dao.User {
	user, _ := r.Context().Value(authUserCtx).(*dao.User)
	return user
}

func checkJWTAuth(w http.ResponseWriter, r *http.Request, app *application.Application) *dao.User {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		responses.Unauthorized(w, r, app.Logger)
		return nil
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) == 0 {
		responses.Unauthorized(w, r, app.Logger)
		return nil
	}

	if len(parts) != 2 || parts[0] != "Bearer" {
		responses.Unauthorized(w, r, app.Logger)
		return nil
	}

	token := parts[1]

	jwtToken, err := app.Authenticator.ValidateToken(token)

	if err != nil {
		app.Logger.Errorw("Error while decoding JWT Token")
		responses.Unauthorized(w, r, app.Logger)
		return nil
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		responses.InternalServerError(w, r, fmt.Errorf("there has been an error while decoding the MapClaims"), app.Logger)
		return nil
	}

	userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)

	if err != nil {
		responses.Unauthorized(w, r, app.Logger)
		return nil
	}

	ctx := r.Context()

	user, err := GetUser(ctx, userID, app)

	if err != nil {
		responses.Unauthorized(w, r, app.Logger)
		return nil
	}

	return user
}
