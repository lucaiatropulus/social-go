package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/internal/mailer"
	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/internal/utils"
	"github.com/lucaiatropulus/social/models"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
	"time"
)

type AuthService struct {
	app *application.Application
}

func NewAuthService(app *application.Application) *AuthService {
	return &AuthService{app}
}

func (s *AuthService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterUserRequest

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		responses.BadRequest(w, r, err, "The register information is not valid")
		return
	}

	if ok := payload.IsValid(); !ok {
		responses.BadRequest(w, r, errors.New("register data invalid"), "The register information is not valid")
		return
	}

	userRole, err := s.app.Store.Roles.GetRoleByName(r.Context(), "user")

	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	user := &dao.User{
		Username: payload.Username,
		Email:    payload.Email,
		RoleID:   userRole.ID,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	uuidToken := uuid.New().String()

	hash := sha256.Sum256([]byte(uuidToken))
	hashedToken := hex.EncodeToString(hash[:])

	err = s.app.Store.Users.CreateAndInvite(ctx, user, hashedToken, s.app.Config.Mail.Expiration)

	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	userWithToken := &models.UserWithToken{
		User:  user,
		Token: uuidToken,
	}

	isProdEnv := s.app.Config.APP.Environment == "production"

	vars := struct {
		Username      string
		ActivationURL string
	}{
		Username:      user.Username,
		ActivationURL: fmt.Sprintf("%s/confirm/%s", s.app.Config.APP.FrontendURL, uuidToken),
	}

	err = s.app.Mailer.Send(mailer.UserWelcomeTemplate, user.Username, user.Email, vars, !isProdEnv)

	if err != nil {

		if err := s.app.Store.Users.Delete(ctx, user.ID); err != nil {
			utils.GetLogger().Errorw("Error deleting user", "error", err)
		}

		responses.InternalServerError(w, r, err)
		return
	}

	if err := responses.JSONResponse(w, http.StatusCreated, userWithToken); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}
}

func (s *AuthService) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload models.LoginRequest

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		responses.BadRequest(w, r, err, "Your credentials could not be identified")
		return
	}

	if ok := payload.IsValid(); !ok {
		responses.BadRequest(w, r, errors.New("error validating login payload"), "Your credentials could not be identified")
		return
	}

	user, err := s.app.Store.Users.GetByEmail(r.Context(), payload.Email)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			responses.Unauthorized(w, r)
		default:
			responses.InternalServerError(w, r, err)
		}
		return
	}

	if ok := user.Password.CheckPassword(payload.Password); !ok {
		responses.Unauthorized(w, r)
		return
	}

	validDuration, ok := utils.ParseStringToDuration(s.app.Config.Auth.ValidDuration)

	if !ok {
		responses.InternalServerError(w, r, errors.New("auth valid duration configuration is malformed"))
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(validDuration).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": s.app.Config.Auth.Issuer,
		"aud": s.app.Config.Auth.Audience,
	}

	token, err := s.app.Authenticator.GenerateToken(claims)

	if err != nil {
		responses.InternalServerError(w, r, err)
		return
	}

	if err := responses.JSONResponse(w, http.StatusOK, token); err != nil {
		responses.InternalServerError(w, r, err)
		return
	}
}

func (s *AuthService) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	err := s.app.Store.Users.Activate(r.Context(), token)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			responses.BadRequest(w, r, err, "We were unable to activate the account")
		default:
			responses.InternalServerError(w, r, err)
		}

		return
	}

	responses.EmptyJSONResponse(w, http.StatusNoContent)
}
