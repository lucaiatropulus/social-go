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

type PostService struct {
	app *application.Application
}

func NewPostService(app *application.Application) *PostService {
	return &PostService{app}
}

func (s *PostService) CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload models.CreatePostRequest

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		responses.BadRequest(w, r, err, "Nu am putut genera", s.app.Logger)
		return
	}

	if reason, isValid := payload.IsValid(); !isValid {
		responses.BadRequest(w, r, nil, reason, s.app.Logger)
		return
	}

	user := middleware.GetAuthenticatedUserFromCtx(r)

	post := models.CreatePostFromPayload(payload, user.ID)

	if err := s.app.Store.Posts.Create(r.Context(), post); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	if err := responses.JSONResponse(w, http.StatusCreated, post); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}
}

func (s *PostService) GetPost(w http.ResponseWriter, r *http.Request) {
	post := middleware.GetPostFromPathParam(r)

	commentsCount, err := s.app.Store.Comments.GetCountByID(r.Context(), post.ID)

	if err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	user, err := s.app.Store.Users.GetByID(r.Context(), post.UserID)

	if err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	postWithComments := &models.DisplayPost{
		Post:          *post,
		Username:      user.Username,
		CommentsCount: commentsCount,
	}

	if err := responses.JSONResponse(w, http.StatusOK, postWithComments); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}
}

func (s *PostService) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := middleware.GetPostFromPathParam(r)

	var payload models.UpdatePost

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		responses.BadRequest(w, r, err, "We were unable to save the changes", s.app.Logger)
		return
	}

	if isValid := payload.IsValid(); !isValid {
		responses.BadRequest(w, r, nil, "You need to change at least one of the fields", s.app.Logger)
		return
	}

	payload.SetUpdatedContent(post)

	if err := s.app.Store.Posts.Update(r.Context(), post); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	if err := responses.JSONResponse(w, http.StatusOK, post); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}
}

func (s *PostService) DeletePost(w http.ResponseWriter, r *http.Request) {
	post := middleware.GetPostFromPathParam(r)

	if err := s.app.Store.Posts.Delete(r.Context(), post.ID); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			responses.NotFound(w, r, err, s.app.Logger)
		default:
			responses.InternalServerError(w, r, err, s.app.Logger)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
