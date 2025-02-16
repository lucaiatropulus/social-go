package services

import (
	"errors"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/middleware"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
)

type FeedService struct {
	app *application.Application
}

func NewFeedService(app *application.Application) *FeedService {
	return &FeedService{app}
}

func (s *FeedService) GetUserFeed(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "DESC",
	}

	fq, err := fq.Parse(r)

	if err != nil {
		responses.BadRequest(w, r, err, "We were unable to deliver the feed", s.app.Logger)
		return
	}

	if ok := fq.IsValid(); !ok {
		responses.BadRequest(w, r, errors.New("nu e ok"), "We were unable to deliver the feed", s.app.Logger)
		return
	}

	user := middleware.GetAuthenticatedUserFromCtx(r)

	feed, err := s.app.Store.Feed.GetUserFeed(r.Context(), user.ID, fq)

	if err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}

	if err := responses.JSONResponse(w, http.StatusOK, feed); err != nil {
		responses.InternalServerError(w, r, err, s.app.Logger)
		return
	}
}
