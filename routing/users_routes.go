package routing

import (
	"github.com/go-chi/chi/v5"
)

func (routing *Routing) configureUsersRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(routing.middleware.JWTAuthMiddleware)
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(routing.middleware.UserIDPathParamMiddleware)

			r.Get("/", routing.handlers.UserHandler.GetUserHandler)
			r.Put("/follow", routing.handlers.UserHandler.FollowUserHandler)
			r.Put("/unfollow", routing.handlers.UserHandler.UnfollowUserHandler)
		})

		r.Group(func(r chi.Router) {
			r.Get("/feed", routing.handlers.FeedHandler.GetUserFeed)
			r.Patch("/update", routing.handlers.UserHandler.UpdateUserHandler)
		})
	})
}
