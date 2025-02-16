package middleware

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/internal/dao"
	"github.com/lucaiatropulus/social/internal/store"
	"github.com/lucaiatropulus/social/responses"
	"net/http"
	"strconv"
)

type postKey string

const postCtx postKey = "post"

func (m *Middleware) PostsPathParamMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		post := checkPostParam(w, r, m.app)

		if post == nil {
			return
		}

		ctx := context.WithValue(r.Context(), postCtx, post)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPostFromPathParam(r *http.Request) *dao.Post {
	post, _ := r.Context().Value(postCtx).(*dao.Post)
	return post
}

func checkPostParam(w http.ResponseWriter, r *http.Request, app *application.Application) *dao.Post {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)

	if err != nil {
		responses.BadRequest(w, r, err, "We were unable to find the post", app.Logger)
		return nil
	}

	post, err := app.Store.Posts.GetByID(r.Context(), postID)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			responses.NotFound(w, r, err, app.Logger)
		default:
			responses.InternalServerError(w, r, err, app.Logger)
		}
		return nil
	}

	return post
}
