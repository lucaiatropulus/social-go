package services

import (
	"errors"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/internal/utils"
	"github.com/lucaiatropulus/social/middleware"
	"github.com/lucaiatropulus/social/models"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
)

type UserService struct {
	app *application.Application
}

func NewUserService(app *application.Application) *UserService {
	return &UserService{app}
}

func (s *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromPathParameter(r)

	user, err := middleware.GetUser(r.Context(), userID, s.app)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			responses.NotFound(w, r, err, s.app.Logger)
		default:
			responses.InternalServerError(w, r, err, s.app.Logger)
		}
		return
	}

	if err := responses.JSONResponse(w, http.StatusOK, user); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
	}
}

func (s *UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetAuthenticatedUserFromCtx(r)
	var payload models.UpdateUserRequest

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		responses.BadRequest(w, r, err, "The info you sent are invalid", s.app.Logger)
		return
	}

	payload.UpdateUser(user)

	if err := s.app.Store.Users.Update(r.Context(), user); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	if s.app.Config.Redis.Enabled {
		_ = s.app.Cache.Users.Delete(r.Context(), user.ID)
	}

	if err := responses.JSONResponse(w, http.StatusOK, user); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}
}

func (s *UserService) FollowUser(w http.ResponseWriter, r *http.Request) {
	followedUserID := middleware.GetUserIDFromPathParameter(r)
	authenticatedUser := middleware.GetAuthenticatedUserFromCtx(r)

	if err := s.app.Store.Followers.Follow(r.Context(), followedUserID, authenticatedUser.ID); err != nil {
		switch {
		case errors.Is(err, store.ErrFollowConflict):
			responses.Conflict(w, r, err, s.app.Logger)
			return
		default:
			responses.InternalServerError(w, r, err, s.app.Logger)
			return
		}
	}

	responses.EmptyJSONResponse(w, http.StatusNoContent)
}

func (s *UserService) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	unfollowedUserID := middleware.GetUserIDFromPathParameter(r)
	authenticatedUser := middleware.GetAuthenticatedUserFromCtx(r)

	if err := s.app.Store.Followers.Unfollow(r.Context(), unfollowedUserID, authenticatedUser.ID); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	responses.EmptyJSONResponse(w, http.StatusNoContent)
}
