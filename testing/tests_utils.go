package testing

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/config"
	"github.com/lucaiatropulus/social/internal/auth"
	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application.Application {
	t.Helper()

	configuration := config.NewMockConfig()
	logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()
	mockCache := cache.NewMockStore()
	authenticator := auth.NewJWTAuthenticator(
		configuration.Auth.Secret,
		configuration.Auth.Audience,
		configuration.Auth.Issuer,
	)

	return &application.Application{
		Config:        *configuration,
		Logger:        logger,
		Store:         mockStore,
		Cache:         mockCache,
		Authenticator: authenticator,
	}
}

func executeRequest(req *http.Request, mux http.Handler, app *application.Application, authenticatedUser *dao.User) *httptest.ResponseRecorder {
	if authenticatedUser != nil {
		validDuration, _ := utils.ParseStringToDuration(app.Config.Auth.ValidDuration)

		claims := jwt.MapClaims{
			"sub": authenticatedUser.ID,
			"exp": time.Now().Add(validDuration).Unix(),
			"iat": time.Now().Unix(),
			"nbf": time.Now().Unix(),
			"iss": app.Config.Auth.Issuer,
			"aud": app.Config.Auth.Audience,
		}

		jwtToken, _ := app.Authenticator.GenerateToken(claims)
		req.Header.Set("Authorization", "Bearer "+jwtToken)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}
