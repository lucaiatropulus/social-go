package handlers

import (
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/services"
)

type Handlers struct {
	AuthHandler   *AuthHandler
	FeedHandler   *FeedHandler
	HealthHandler *HealthHandler
	PostHandler   *PostHandler
	UserHandler   *UserHandler
}

func NewHandlers(app *application.Application) *Handlers {
	authService := services.NewAuthService(app)
	feedService := services.NewFeedService(app)
	healthService := services.NewHealthService(app)
	postService := services.NewPostService(app)
	userService := services.NewUserService(app)

	return &Handlers{
		AuthHandler:   &AuthHandler{authService},
		FeedHandler:   &FeedHandler{feedService},
		HealthHandler: &HealthHandler{healthService},
		PostHandler:   &PostHandler{postService},
		UserHandler:   &UserHandler{userService},
	}
}
